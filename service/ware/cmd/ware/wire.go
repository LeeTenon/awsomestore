//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "awsomestore/service/ware/internal/biz"
    "awsomestore/service/ware/internal/conf"
    "awsomestore/service/ware/internal/data"
    "awsomestore/service/ware/internal/server"
    "awsomestore/service/ware/internal/service"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
