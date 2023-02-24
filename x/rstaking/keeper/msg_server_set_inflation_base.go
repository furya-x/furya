package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oldfurya/furya/x/rstaking/types"
	sudotypes "github.com/oldfurya/furya/x/sudo/types"
)

func (k msgServer) SetInflationBase(goCtx context.Context, msg *types.MsgSetInflationBase) (*types.MsgSetInflationBaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetInflationBase(ctx, msg.InflationBase)

	return &types.MsgSetInflationBaseResponse{}, nil
}
