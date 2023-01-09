package keeper

import (
	"github.com/warmage-sports/katana/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
