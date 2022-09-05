//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "awsomestore/access/gate/internal/router"
    "github.com/go-kratos/kratos/v2"
    "github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
    panic(wire.Build(router.ProviderSet, newGinService))
}
