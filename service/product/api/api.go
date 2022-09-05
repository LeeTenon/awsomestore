package api

import (
    "awsomestore/api/product"
    "awsomestore/common/discover"
    "github.com/go-kratos/kratos/v2/log"
)

type ProductClient struct {
    product.ProductClient
}

func NewProductClient(c *discover.Config) *ProductClient {
    conn, err := discover.Discover(c)
    if err != nil {
        log.Fatal(err)
        return nil
    }

    return &ProductClient{
        product.NewProductClient(conn),
    }
}
