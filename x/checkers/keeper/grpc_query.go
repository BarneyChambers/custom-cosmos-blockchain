package keeper

import (
	"github.com/barneychambers/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
