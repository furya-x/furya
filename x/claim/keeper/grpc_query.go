package keeper

import (
	"github.com/oldfurya/furya/x/claim/types"
)

var _ types.QueryServer = Keeper{}
