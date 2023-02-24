package keeper

import (
	"github.com/oldfurya/furya/x/rstaking/types"
)

var _ types.QueryServer = Keeper{}
