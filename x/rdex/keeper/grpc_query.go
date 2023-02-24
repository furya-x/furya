package keeper

import (
	"github.com/oldfurya/furya/x/rdex/types"
)

var _ types.QueryServer = Keeper{}
