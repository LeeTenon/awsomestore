package api

import (
    "awsomestore/api/user"
    "awsomestore/common/discover"
    "github.com/go-kratos/kratos/v2/log"
    "google.golang.org/grpc"
)

type UserClient struct {
    conn    *grpc.ClientConn
    Account user.AccountClient
    Cart    user.CartClient
}

func NewUserClient(c *discover.Config) *UserClient {
    conn, err := discover.Discover(c)
    if err != nil {
        log.Fatal(err)
        return nil
    }

    return &UserClient{
        conn:    conn,
        Account: user.NewAccountClient(conn),
        Cart:    user.NewCartClient(conn),
    }
}