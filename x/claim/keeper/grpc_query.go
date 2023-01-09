package keeper

import (
	"github.com/warmage-sports/katana/x/claim/types"
)

var _ types.QueryServer = Keeper{}
