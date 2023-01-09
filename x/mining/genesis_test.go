package mining_test

import (
	"testing"

	keepertest "github.com/warmage-sports/katana/testutil/keeper"
	"github.com/warmage-sports/katana/testutil/nullify"
	"github.com/warmage-sports/katana/x/mining"
	"github.com/warmage-sports/katana/x/mining/types"
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
