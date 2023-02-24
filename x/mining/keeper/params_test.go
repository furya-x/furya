package keeper_test

import (
	"testing"

	testkeeper "github.com/oldfurya/furya/testutil/keeper"
	"github.com/oldfurya/furya/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MiningKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
