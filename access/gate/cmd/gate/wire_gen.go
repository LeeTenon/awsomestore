// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"awsomestore/access/gate/internal/biz"
	"awsomestore/access/gate/internal/conf"
	"awsomestore/access/gate/internal/data"
	"awsomestore/access/gate/internal/server"
	"awsomestore/access/gate/internal/service"
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
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
