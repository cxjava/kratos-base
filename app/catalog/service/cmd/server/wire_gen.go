// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"kratos-base/app/catalog/service/internal/biz"
	"kratos-base/app/catalog/service/internal/conf"
	"kratos-base/app/catalog/service/internal/data"
	"kratos-base/app/catalog/service/internal/server"
	"kratos-base/app/catalog/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client, err := data.NewEntClient(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, client, logger)
	if err != nil {
		return nil, nil, err
	}
	beerRepo := data.NewBeerRepo(dataData, logger)
	catalogUseCase := biz.NewBeerUseCase(beerRepo, logger)
	catalogService := service.NewCatalogService(catalogUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, catalogService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
