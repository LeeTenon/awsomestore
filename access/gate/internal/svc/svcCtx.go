package svc

import (
    "awsomestore/common/discover"
    productApi "awsomestore/service/product/api"
    userApi "awsomestore/service/user/api"
)

type ServiceContext struct {
    User    *userApi.UserClient
    Product *productApi.ProductClient
}

func NewServiceContext() *ServiceContext {
    return &ServiceContext{
        User: userApi.NewUserClient(&discover.Config{
            Hosts:     []string{"127.0.0.1:2379"},
            Service:   "user",
            Namespace: "store-server",
        }),
        Product: productApi.NewProductClient(&discover.Config{
            Hosts:     []string{"127.0.0.1:2379"},
            Service:   "product",
            Namespace: "store-server",
        }),
    }
}
