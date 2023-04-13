package keeper

import (
	"github.com/furyahub/furyahub/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
