package keeper

import (
	"github.com/enflow.io/enf2/x/enf2/types"
)

var _ types.QueryServer = Keeper{}
