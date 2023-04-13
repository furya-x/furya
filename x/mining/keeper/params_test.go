package keeper_test

import (
	"testing"

	testkeeper "github.com/furyahub/furyahub/testutil/keeper"
	"github.com/furyahub/furyahub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MiningKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
