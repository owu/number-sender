package service

import (
	"github.com/google/wire"
	"number-sender/internal/pkg/calculate"
	"number-sender/internal/pkg/config"
	"number-sender/internal/pkg/limit"
	"number-sender/internal/pkg/redis"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	redis.ProviderSet,
	calculate.ProviderSet,
	limit.ProviderSet,
)
