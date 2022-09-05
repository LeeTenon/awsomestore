package discover

import (
    "context"
    "fmt"
    "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
    kgpc "github.com/go-kratos/kratos/v2/transport/grpc"
    clientv3 "go.etcd.io/etcd/client/v3"
    "google.golang.org/grpc"
)

type Config struct {
    Hosts     []string
    Service   string
    Namespace string
}

func Discover(c *Config) (*grpc.ClientConn, error) {
    client, err := clientv3.New(clientv3.Config{
        Endpoints: c.Hosts,
    })
    if err != nil {
        return nil, err
    }

    discover := etcd.New(client, etcd.Namespace(c.Namespace))

    endpoint := fmt.Sprintf("discovery:///%s", c.Service)
    conn, err := kgpc.DialInsecure(context.Background(), kgpc.WithEndpoint(endpoint), kgpc.WithDiscovery(discover))
    if err != nil {
        return nil, err
    }

    return conn, nil
}
