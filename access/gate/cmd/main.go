package main

import (
    "awsomestore/access/gate/internal/router"
    "awsomestore/access/gate/internal/svc"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := newGinService()

    log.Fatal(r.Run(":7999"))
}

func newGinService() *gin.Engine {
    r := gin.Default()
    router.RegisterRouter(svc.NewServiceContext(), r)

    return r
}
