package keeper

import (
	"github.com/furyahub/furyahub/x/claim/types"
)

var _ types.QueryServer = Keeper{}
