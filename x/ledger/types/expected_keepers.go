package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SudoKeeper interface {
	// Methods imported from sudo should be defined here
	IsAdmin(ctx sdk.Context, address string) bool
}

type RateKeeper interface {
	TokenToRtoken(ctx sdk.Context, denom string, balance sdk.Int) sdk.Int
	GetRate(ctx sdk.Context, denom string) *sdk.Dec
	SetRate(ctx sdk.Context, denom string, total, rtotal sdk.Int) sdk.Dec
}

// BankKeeper defines the contract needed to be fulfilled for banking and supply
// dependencies.
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}
