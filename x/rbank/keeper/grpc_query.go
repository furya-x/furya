package keeper

import (
	"github.com/oldfurya/furya/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
