package redis

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDefaultRedis,
)
