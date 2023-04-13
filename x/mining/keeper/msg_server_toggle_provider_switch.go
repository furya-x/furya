package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furyahub/x/mining/types"
	sudotypes "github.com/furyahub/furyahub/x/sudo/types"
)

func (k msgServer) ToggleProviderSwitch(goCtx context.Context, msg *types.MsgToggleProviderSwitch) (*types.MsgToggleProviderSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.ToggleMiningProviderSwitch(ctx)

	return &types.MsgToggleProviderSwitchResponse{}, nil
}
