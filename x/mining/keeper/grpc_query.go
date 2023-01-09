package keeper

import (
	"github.com/warmage-sports/katana/x/mining/types"
)

var _ types.QueryServer = Keeper{}
