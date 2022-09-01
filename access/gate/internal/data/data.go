package data

import (
    "awsomestore/access/gate/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
    // TODO wrapped database client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    return &Data{}, cleanup, nil
}
