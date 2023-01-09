package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/warmage-sports/katana/testutil/keeper"
	"github.com/warmage-sports/katana/x/rvalidator/keeper"
	"github.com/warmage-sports/katana/x/rvalidator/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RvalidatorKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
