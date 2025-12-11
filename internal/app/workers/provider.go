package workers

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewWorkers,
)
