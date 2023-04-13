package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furyahub/x/mining/types"
	sudotypes "github.com/furyahub/furyahub/x/sudo/types"
)

func (k msgServer) RmRewardToken(goCtx context.Context, msg *types.MsgRmRewardToken) (*types.MsgRmRewardTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.RemoveRewardToken(ctx, msg.Denom)

	return &types.MsgRmRewardTokenResponse{}, nil
}
