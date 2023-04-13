package keeper

import (
	"github.com/furyahub/furyahub/x/rdex/types"
)

var _ types.QueryServer = Keeper{}
