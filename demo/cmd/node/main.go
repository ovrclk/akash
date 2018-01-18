package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"

	sdk "github.com/cosmos/cosmos-sdk"
	client "github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/cosmos/cosmos-sdk/modules/auth"
	"github.com/cosmos/cosmos-sdk/modules/base"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	_ "github.com/cosmos/cosmos-sdk/modules/eyes"
	_ "github.com/cosmos/cosmos-sdk/modules/fee"
	"github.com/cosmos/cosmos-sdk/modules/ibc"
	"github.com/cosmos/cosmos-sdk/modules/nonce"
	"github.com/cosmos/cosmos-sdk/modules/roles"
	"github.com/cosmos/cosmos-sdk/server/commands"
	"github.com/cosmos/cosmos-sdk/stack"
	_ "github.com/ovrclk/photon/demo/plugins/accounts"
)

// RootCmd is the entry point for this binary
var RootCmd = &cobra.Command{
	Use:   "photon",
	Short: "Blockchained infrustructure",
}

// BuildApp constructs the stack we want to use for this app
func BuildApp(feeDenom string) sdk.Handler {
	return stack.New(
		base.Logger{},
		stack.Recovery{},
		auth.Signatures{},
		base.Chain{},
		stack.Checkpoint{OnCheck: true},
		nonce.ReplayCheck{},
	).
		IBC(ibc.NewMiddleware()).
		Apps(
			roles.NewMiddleware(),
			// fee.NewSimpleFeeMiddleware(coin.Coin{feeDenom, 0}, fee.Bank),
			stack.Checkpoint{OnDeliver: true},
		).
		Dispatch(
			coin.NewHandler(),
			/*
				add module/plugin handlers here
				use a stack.WrapHandler over all handlers after coin.NewHandler()
			*/
			// stack.WrapHandler(accounts.NewHandler()),
			// stack.WrapHandler(roles.NewHandler()),
			// stack.WrapHandler(ibc.NewHandler()),
			// and just for run, add eyes as well
			// stack.WrapHandler(eyes.NewHandler()),
		)
}

func main() {
	// require all fees in mycoin - change this in your app!
	commands.Handler = BuildApp("photon")

	RootCmd.AddCommand(
		commands.InitCmd,
		commands.StartCmd,
		//commands.RelayCmd,
		commands.UnsafeResetAllCmd,
		client.VersionCmd,
	)
	commands.SetUpRoot(RootCmd)

	cmd := cli.PrepareMainCmd(RootCmd, "BC", os.ExpandEnv("$HOME/.demonode"))
	cmd.Execute()
}
