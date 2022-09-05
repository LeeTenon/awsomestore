package controller

import (
    "awsomestore/access/gate/internal/service"
    "awsomestore/access/gate/internal/svc"
    "awsomestore/api/gate"
    "context"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

type UserController struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUserController(svcCtx *svc.ServiceContext) *UserController {
    return &UserController{
        ctx:    context.Background(),
        svcCtx: svcCtx,
    }
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    req := &gate.CreateUserReq{}
    if err := ctx.ShouldBindJSON(req); err != nil {
        log.Println(err)
    }

    s := service.NewUserService(c.ctx, c.svcCtx)
    if err := s.CreateUser(req); err != nil {
        ctx.JSON(http.StatusInternalServerError, gate.ErrorInfo{
            Code:   http.StatusInternalServerError,
            Reason: err.Error(),
            Msg:    "",
        })
        return
    }

    ctx.JSON(http.StatusOK, gate.CreateUserResp{})
}

func (c *UserController) QueryCart(ctx *gin.Context) {
    req := &gate.QueryCartReq{}
    if err := ctx.ShouldBindQuery(req); err != nil {
        log.Println(err)
    }

    s := service.NewUserService(c.ctx, c.svcCtx)
    resp, err := s.QueryCart(req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gate.ErrorInfo{
            Code:   http.StatusInternalServerError,
            Reason: err.Error(),
            Msg:    "",
        })
        return
    }

    ctx.JSON(http.StatusOK, resp)
}
