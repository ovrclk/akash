package main

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/gosuri/uitable"
	"github.com/ovrclk/akash/cmd/akash/session"
	"github.com/ovrclk/akash/cmd/common"
	"github.com/ovrclk/akash/errors"
	"github.com/ovrclk/akash/util/uiutil"
	"github.com/spf13/cobra"

	. "github.com/ovrclk/akash/util"
)

func keyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Manage keys",
	}
	cmd.AddCommand(keyCreateCommand())
	cmd.AddCommand(keyListCommand())
	cmd.AddCommand(keyShowCommand())
	cmd.AddCommand(keyRecoverCommand())
	return cmd
}

func keyCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <name>",
		Short:   "create a new key locally",
		Long:    "Create a new key and store it locally.\nTo recover a key using the recovery codes generated by the command, see 'akash help key recover'",
		Example: keyCreateHelp,
		RunE:    session.WithSession(session.RequireRootDir(doKeyCreateCommand)),
	}
	session.AddFlagKeyType(cmd, cmd.Flags())
	return cmd
}

func doKeyCreateCommand(ses session.Session, cmd *cobra.Command, args []string) error {
	var name string
	if len(args) == 1 {
		name = args[0]
	}
	name = ses.Mode().Ask().StringVar(name, "Key Name (required): ", true)
	if len(name) == 0 {
		return fmt.Errorf("required argument missing: name")
	}

	kmgr, err := ses.KeyManager()
	if err != nil {
		return err
	}

	ktype, err := ses.KeyType()
	if err != nil {
		return err
	}

	password, err := ses.Password()
	if err != nil {
		return err
	}

	info, seed, err := kmgr.CreateMnemonic(name, common.DefaultCodec, password, ktype)
	if err != nil {
		return err
	}

	pdata := session.NewPrinterDataKV().
		AddResultKV("name", name).
		AddResultKV("public_key_address", X(info.GetPubKey().Address())).
		AddResultKV("recovery_seed", seed)

	return ses.Mode().
		When(session.ModeTypeText, func() error {
			return session.NewTextPrinter(pdata, nil).Flush()
		}).
		When(session.ModeTypeJSON, func() error {
			return session.NewJSONPrinter(pdata, nil).Flush()
		}).
		When(session.ModeTypeInteractive, func() error {
			fmt.Println(string(uiutil.NewTitle(fmt.Sprintf("Successfully created key for '%s'", name)).Bytes()))
			table := uitable.New()
			table.AddRow("Public Key:", X(info.GetPubKey().Address()))
			table.AddRow("Recovery Codes:", seed)
			fmt.Println(table)
			return nil
		}).Run()
}

func keyListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all the keys stored locally",
		RunE:  session.WithSession(session.RequireKeyManager(doKeyListCommand)),
	}
}

func doKeyListCommand(s session.Session, cmd *cobra.Command, args []string) error {
	kmgr, _ := s.KeyManager()
	infos, err := kmgr.List()
	if err != nil {
		return err
	}
	pdata := session.NewPrinterDataList()
	for _, info := range infos {
		d := map[string]string{
			"name":               info.GetName(),
			"public_key_address": X(info.GetPubKey().Address()),
		}
		pdata.AddResultList(d)
	}
	pdata.Raw = infos

	s.Mode().When(session.ModeTypeInteractive, func() error {
		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true

		table.AddRow(
			uiutil.NewTitle("Name").String(),
			uiutil.NewTitle("Public Key (Address)").String(),
		)

		for _, info := range pdata.Result {
			table.AddRow(info["name"], info["public_key_address"])
		}
		fmt.Println(table)
		return nil
	}).When(session.ModeTypeText, func() error {
		return session.NewTextPrinter(pdata, nil).Flush()
	}).When(session.ModeTypeJSON, func() error {
		return session.NewJSONPrinter(pdata, nil).Flush()
	})
	return s.Mode().Run()
}

func keyRecoverCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "recover <name> <recovery-codes>...",
		Short:   "recover a key using recovery codes",
		Long:    "Recover a key using the recovery code generated during key creation and store it locally. For help with creating a key, see 'akash help key create'",
		Example: keyRecoverExample,
		RunE:    session.WithSession(session.RequireKeyManager(doKeyRecoverCommand)),
	}
}

func doKeyRecoverCommand(ses session.Session, cmd *cobra.Command, args []string) error {
	// the first arg is the key name and the rest are mnemonic codes
	var name, seed string

	if len(args) > 2 {
		name, args = args[0], args[1:]
		seed = strings.Join(args, " ")
	}

	name = ses.Mode().Ask().StringVar(name, "Key Name (required): ", true)
	seed = ses.Mode().Ask().StringVar(seed, "Recovery Codes (required): ", true)
	if len(name) == 0 {
		return errors.NewArgumentError("required argument missing: name")
	}
	if len(seed) == 0 {
		return errors.NewArgumentError("seed")
	}

	password, err := ses.Password()
	if err != nil {
		return err
	}
	kmgr, _ := ses.KeyManager()
	params := *hd.NewFundraiserParams(0, 0)
	info, err := kmgr.Derive(name, seed, keys.DefaultBIP39Passphrase, password, params)
	if err != nil {
		return err
	}

	pdata := session.NewPrinterDataKV().
		AddResultKV("name", name).
		AddResultKV("public_key_address", X(info.GetPubKey().Address()))
	pdata.Raw = info

	return ses.Mode().
		When(session.ModeTypeText, func() error {
			return session.NewTextPrinter(pdata, nil).Flush()
		}).
		When(session.ModeTypeJSON, func() error {
			return session.NewJSONPrinter(pdata, nil).Flush()
		}).
		When(session.ModeTypeInteractive, func() error {
			title := uiutil.NewTitle(fmt.Sprintf("Successfully recovered key, stored locally as '%s'", name))
			fmt.Println(string(title.Bytes()))
			table := uitable.New()
			table.AddRow("Public Key:", X(info.GetPubKey().Address()))
			fmt.Println(table)
			return nil
		}).
		Run()
}

func keyShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <name>",
		Short: "display a key",
		Args:  cobra.ExactArgs(1),
		RunE:  session.WithSession(session.RequireRootDir(doKeyShowCommand)),
	}
	session.AddFlagKeyType(cmd, cmd.Flags())
	return cmd
}

func doKeyShowCommand(ses session.Session, cmd *cobra.Command, args []string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	}

	if len(name) == 0 {
		return errors.NewArgumentError("name")
	}

	kmgr, err := ses.KeyManager()
	if err != nil {
		return err
	}
	info, err := kmgr.Get(name)
	if err != nil {
		return err
	}

	if len(info.GetPubKey().Address()) == 0 {
		return fmt.Errorf("key not found %s", name)
	}
	pdata := session.NewPrinterDataKV().
		AddResultKV("name", name).
		AddResultKV("public_key_address", X(info.GetPubKey().Address()))

	return ses.Mode().
		When(session.ModeTypeInteractive, func() error {
			fmt.Println(pdata.Result[0]["public_key_address"])
			return nil
		}).
		When(session.ModeTypeText, func() error {
			return session.NewTextPrinter(pdata, nil).Flush()
		}).
		When(session.ModeTypeJSON, func() error {
			return session.NewJSONPrinter(pdata, nil).Flush()
		}).
		Run()
}

var (
	keyCreateHelp = `
- Create a key with the name 'greg':

  $ akash key create greg

  Successfully created key for 'greg'
  ===================================

  Public Key:    	f4e03226c054b1adafaa2739bad720c095500a49
  Recovery Codes:	figure share industry canal...
`

	keyRecoverExample = `
- Recover a key with the name 'my-key':

  $ akash key recover my-key today napkin arch \
	 picnic fox case thrive table journey ill  \
	 any enforce awesome desert chapter regret \
	 narrow capable advice skull pipe giraffe \
	 clown outside
`
)
