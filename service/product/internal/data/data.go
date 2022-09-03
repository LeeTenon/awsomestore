package data

import (
    "awsomestore/service/product/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)

// Data .
type Data struct {
    DB    *gorm.DB
    Redis string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }

    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:                      c.Database.Source,
        DisableDatetimePrecision: true,
        DontSupportRenameIndex:   true,
    }), &gorm.Config{})
    if err != nil {
        log.Fatalf("init db error: %s", err.Error())
    }

    return &Data{
        DB: db,
    }, cleanup, nil
}
