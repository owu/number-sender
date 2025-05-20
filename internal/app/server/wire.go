//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/owu/number-sender/internal/app/service"
	"github.com/owu/number-sender/internal/app/workers"
)

func InitApp() error {
	panic(wire.Build(
		wire.Struct(new(Options), "*"),
		service.ProviderSet,
		workers.ProviderSet,
		initApp,
	))

	return nil
}
