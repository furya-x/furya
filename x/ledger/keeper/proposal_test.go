package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furyahub/furyahub/testutil/keeper"
	"github.com/furyahub/furyahub/x/ledger/keeper"
	//"github.com/furyahub/furyahub/x/ledger/types"
)

func setupSettings(t testing.TB) {
	k, ctx := keepertest.LedgerKeeper(t)
	s, _ := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
	t.Log(s)
}
