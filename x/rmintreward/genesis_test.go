package rmintreward_test

import (
	"testing"

	keepertest "github.com/furyahub/furyahub/testutil/keeper"
	"github.com/furyahub/furyahub/testutil/nullify"
	"github.com/furyahub/furyahub/x/rmintreward"
	"github.com/furyahub/furyahub/x/rmintreward/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RmintrewardKeeper(t)
	rmintreward.InitGenesis(ctx, *k, genesisState)
	got := rmintreward.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
