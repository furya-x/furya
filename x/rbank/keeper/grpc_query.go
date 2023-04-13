package keeper

import (
	"github.com/furyahub/furyahub/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
