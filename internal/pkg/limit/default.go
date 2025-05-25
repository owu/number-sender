package limit

import (
	"github.com/owu/number-sender/internal/pkg/config"
	"sync"
)

type Options struct {
	Configs *config.LoadConfigs
}

type Limiter struct {
	Options
	BucketMap sync.Map // map[string]*ratelimit.Bucket
}

func NewLimiter(options Options) *Limiter {
	return &Limiter{
		Options:   options,
		BucketMap: sync.Map{},
	}
}
