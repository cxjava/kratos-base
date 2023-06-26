//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-base/app/catalog/service/internal/biz"
	"kratos-base/app/catalog/service/internal/conf"
	"kratos-base/app/catalog/service/internal/data"
	"kratos-base/app/catalog/service/internal/server"
	"kratos-base/app/catalog/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
