package router

import (
    "awsomestore/access/gate/internal/controller"
    "awsomestore/access/gate/internal/svc"
    "github.com/gin-gonic/gin"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(RegisterRouter)

func RegisterRouter(s *svc.ServiceContext, r *gin.Engine) {
    user := controller.NewUserController(s)
    product := controller.NewProductController(s)

    r.POST("/login")
    r.POST("/register", user.CreateUser)
    v1 := r.Group("/v1")
    {
        // 购物车
        v1.GET("/cart/info", user.QueryCart)
        v1.POST("/cart/update")
        // 商品
        v1.PUT("/product")
        v1.POST("/product", product.CreateProduct)
        v1.GET("/product/popular", product.QueryPopularProduct)
        v1.GET("/product/detail")
    }
}
