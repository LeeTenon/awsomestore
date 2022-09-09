package main

import (
	"awsomestore/common/logx"
	"awsomestore/service/user/internal/conf"
	"flag"
	"os"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
    // Name is the name of the compiled software.
    Name string = "user"
    // Version is the version of the compiled software.
    Version string
    // flagconf is the config flag.
    flagconf string

    id, _ = os.Hostname()
)

func init() {
    flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server) *kratos.App {
    // new etcd client
    client, err := clientv3.New(clientv3.Config{
        Endpoints: []string{"127.0.0.1:2379"},
    })
    if err != nil {
        panic(err)
    }
    // new reg with etcd client
    reg := etcd.New(client, etcd.Namespace("store-server"))
    return kratos.New(
        kratos.ID(id),
        kratos.Name(Name),
        kratos.Version(Version),
        kratos.Metadata(map[string]string{}),
        kratos.Logger(logger),
        kratos.Server(
            gs,
        ),
        kratos.Registrar(reg),
    )
}

func main() {
    flag.Parse()
    logger := logx.New(Name, os.Stdout)
    c := config.New(
        config.WithSource(
            file.NewSource(flagconf),
        ),
    )
    defer c.Close()

    if err := c.Load(); err != nil {
        panic(err)
    }

    var bc conf.Bootstrap
    if err := c.Scan(&bc); err != nil {
        panic(err)
    }

    app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
    if err != nil {
        panic(err)
    }
    defer cleanup()

    // start and wait for stop signal
    if err = app.Run(); err != nil {
        panic(err)
    }
}
