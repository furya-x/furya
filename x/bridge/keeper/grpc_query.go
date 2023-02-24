package keeper

import (
	"github.com/oldfurya/furya/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
