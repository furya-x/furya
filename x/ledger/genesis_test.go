package ledger_test

import (
	"testing"

	keepertest "github.com/oldfurya/furya/testutil/keeper"
	"github.com/oldfurya/furya/x/ledger"
	"github.com/oldfurya/furya/x/ledger/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LedgerKeeper(t)
	ledger.InitGenesis(ctx, k, genesisState)
	got := ledger.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
