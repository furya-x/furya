package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furyahub/x/rbank/types"
	sudoTypes "github.com/furyahub/furyahub/x/sudo/types"
)

func (k msgServer) AddDenom(goCtx context.Context, msg *types.MsgAddDenom) (*types.MsgAddDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.SetAccAddressPrefix(ctx, msg.Metadata.Base, msg.AccAddressPrefix)
	k.SetValAddressPrefix(ctx, msg.Metadata.Base, msg.ValAddressPrefix)
	k.bankKeeper.SetDenomMetaData(ctx, msg.Metadata)

	return &types.MsgAddDenomResponse{}, nil
}
