package cmd

import (
	"crypto/tls"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	akashclient "github.com/ovrclk/akash/client"
	cmdcommon "github.com/ovrclk/akash/cmd/common"
	gwrest "github.com/ovrclk/akash/provider/gateway/rest"
	cutils "github.com/ovrclk/akash/x/cert/utils"
	mcli "github.com/ovrclk/akash/x/market/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

const (
	FlagStdin = "stdin"
	FlagTty = "tty"
)

func leaseShellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.MinimumNArgs(1),
		Use:          "lease-shell",
		Short:        "do lease shell",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLeaseShell(cmd, args)
		},
	}

	addLeaseFlags(cmd)
	cmd.Flags().Bool(FlagStdin, false, "connect stdin")
	if err := viper.BindPFlag(FlagStdin, cmd.Flags().Lookup(FlagStdin)); err != nil {
		return nil
	}

	cmd.Flags().Bool(FlagTty, false, "connect an interactive terminal")
	if err := viper.BindPFlag(FlagTty, cmd.Flags().Lookup(FlagTty)); err != nil {
		return nil
	}



	return cmd
}

func doLeaseShell(cmd *cobra.Command, args []string) error {
	var stdin io.Reader
	connectStdin := viper.GetBool(FlagStdin)
	if connectStdin {
		stdin = os.Stdin
	}

	cctx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	prov, err := providerFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	bid, err := mcli.BidIDFromFlagsForOwner(cmd.Flags(), cctx.FromAddress)
	if err != nil {
		return err
	}

	cert, err := cutils.LoadAndQueryCertificateForAccount(cmd.Context(), cctx, cctx.Keyring)
	if err != nil {
		return err
	}

	gclient, err := gwrest.NewClient(akashclient.NewQueryClientFromCtx(cctx), prov, []tls.Certificate{cert})
	if err != nil {
		return err
	}

	service := args[0]
	remoteCmd := args[1:]
	err = gclient.LeaseShell(cmd.Context(), bid.LeaseID(), service, remoteCmd, stdin, os.Stdout, os.Stderr)
	if err != nil {
		return showErrorToUser(err)
	}

	return cmdcommon.PrintJSON(cctx, "done")
}

