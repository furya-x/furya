package keeper

import (
	"github.com/furyahub/furyahub/x/mining/types"
)

var _ types.QueryServer = Keeper{}
