package limit

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewLimiter,
	wire.Struct(new(Options), "*"),
)
