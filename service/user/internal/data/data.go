package data

import (
    "awsomestore/service/user/internal/conf"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

var ProviderSet = wire.NewSet(NewData, NewAccountRepo, NewCartRepo)

type Data struct {
    DB      *gorm.DB
    MongoDB mongo.Database
    Redis   string
}

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

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Database.Source).SetConnectTimeout(5*time.Second))
    mongoDB = client.Database("")

    return &Data{
        DB: db,
    }, cleanup, nil
}
