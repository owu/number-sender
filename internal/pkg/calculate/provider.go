package calculate

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewChains,
)
