package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ovrclk/akash/x/provider/types"
)

const (
	keyProvider  = "Provider"
	keyProviders = "Providers"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyProvider,
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenValidator(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyProviders,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenValidators(r))
			},
		),
	}
}
