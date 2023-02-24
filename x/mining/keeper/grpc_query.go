package keeper

import (
	"github.com/oldfurya/furya/x/mining/types"
)

var _ types.QueryServer = Keeper{}
