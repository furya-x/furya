package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/bridge/types"
	sudoTypes "github.com/warmage-sports/katana/x/sudo/types"
)

func (k msgServer) RmBannedDenom(goCtx context.Context, msg *types.MsgRmBannedDenom) (*types.MsgRmBannedDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	k.Keeper.RmBannedDenom(ctx, uint8(msg.ChainId), msg.Denom)

	return &types.MsgRmBannedDenomResponse{}, nil
}
