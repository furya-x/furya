package mining_test

import (
	"testing"

	keepertest "github.com/furyahub/furyahub/testutil/keeper"
	"github.com/furyahub/furyahub/testutil/nullify"
	"github.com/furyahub/furyahub/x/mining"
	"github.com/furyahub/furyahub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.MiningKeeper(t)
	mining.InitGenesis(ctx, *k, *genesisState)
	got := mining.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
