package rvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warmage-sports/katana/x/rvalidator/keeper"
	"github.com/warmage-sports/katana/x/rvalidator/types"
	rvotetypes "github.com/warmage-sports/katana/x/rvote/types"
)

func NewProposalHandler(k keeper.Keeper) rvotetypes.Handler {
	return func(ctx sdk.Context, content rvotetypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateRValidatorProposal:
			return k.ProcessUpdateRValidatorProposal(ctx, c)
		case *types.UpdateRValidatorReportProposal:
			return k.ProcessUpdateRValidatorReportProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
