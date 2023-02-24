package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oldfurya/furya/x/rdex/types"
	sudotypes "github.com/oldfurya/furya/x/sudo/types"
)

func (k msgServer) ToggleProviderSwitch(goCtx context.Context, msg *types.MsgToggleProviderSwitch) (*types.MsgToggleProviderSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.ToggleProviderSwitch(ctx)
	return &types.MsgToggleProviderSwitchResponse{}, nil
}
