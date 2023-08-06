//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"taa-ex/internal/biz"
	"taa-ex/internal/conf"
	"taa-ex/internal/data"
	"taa-ex/internal/server"
	"taa-ex/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/go-kratos/kratos/v2"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
