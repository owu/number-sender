package service

import (
	"github.com/google/wire"
	"github.com/owu/number-sender/internal/pkg/calculate"
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/limit"
	"github.com/owu/number-sender/internal/pkg/redis"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	redis.ProviderSet,
	calculate.ProviderSet,
	limit.ProviderSet,
)
