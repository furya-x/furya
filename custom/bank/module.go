package bank

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/warmage-sports/katana/utils"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the staking module.
type AppModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultGenesisState := types.DefaultGenesisState()
	// {
	//     "description": "The native staking token of the Katana Hub",
	//     "denom_units": [
	//         {
	//             "denom": "ukata",
	//             "exponent": 0,
	//             "aliases": [
	//                 "microkata"
	//             ]
	//         },
	//         {
	//             "denom": "mkata",
	//             "exponent": 3,
	//             "aliases": [
	//               "millikata"
	//             ]
	//         },
	//         {
	//             "denom": "kata",
	//             "exponent": 6
	//         }
	//     ],
	//     "base": "ukata",
	//     "display": "kata",
	//     "name": "KATA",
	//     "symbol": "KATA"
	// }
	defaultGenesisState.DenomMetadata = append(defaultGenesisState.DenomMetadata,
		types.Metadata{
			Description: "The native staking token of the Katana Hub",
			DenomUnits: []*types.DenomUnit{
				{
					Denom:    utils.KataDenom,
					Exponent: 0,
					Aliases:  []string{"microkata"},
				},
				{
					Denom:    "mkata",
					Exponent: 3,
					Aliases:  []string{"millikata"},
				},
				{
					Denom:    "kata",
					Exponent: 6,
				},
			},
			Base:    utils.KataDenom,
			Display: "kata",
			Name:    "KATA",
			Symbol:  "KATA",
		},
	)
	return cdc.MustMarshalJSON(defaultGenesisState)
}
