package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furyahub/furyahub/testutil/keeper"
	"github.com/furyahub/furyahub/x/ledger/keeper"
	"github.com/furyahub/furyahub/x/ledger/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LedgerKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
