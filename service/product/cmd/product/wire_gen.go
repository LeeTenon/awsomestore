// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
    "awsomestore/service/product/internal/biz"
    "awsomestore/service/product/internal/conf"
    "awsomestore/service/product/internal/data"
    "awsomestore/service/product/internal/server"
    "awsomestore/service/product/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	productRepo := data.NewProductRepo(dataData, logger)
	productUsecase := biz.NewProductUsecase(productRepo, logger)
	productService := service.NewProductService(productUsecase)
	grpcServer := server.NewGRPCServer(confServer, productService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
