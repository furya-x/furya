package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/warmage-sports/katana/testutil/keeper"
	"github.com/warmage-sports/katana/x/rstaking/keeper"
	"github.com/warmage-sports/katana/x/rstaking/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RStakingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
