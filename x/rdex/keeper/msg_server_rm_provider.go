package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furyahub/x/rdex/types"
	sudotypes "github.com/furyahub/furyahub/x/sudo/types"
)

func (k msgServer) RmProvider(goCtx context.Context, msg *types.MsgRmProvider) (*types.MsgRmProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	provider, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return nil, err
	}
	k.Keeper.RemoveProvider(ctx, provider)

	return &types.MsgRmProviderResponse{}, nil
}
