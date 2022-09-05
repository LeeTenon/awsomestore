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

type ProductController struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewProductController(svcCtx *svc.ServiceContext) *ProductController {
    return &ProductController{
        ctx:    context.Background(),
        svcCtx: svcCtx,
    }
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    req := &gate.CreateProductReq{}
    if err := ctx.ShouldBindJSON(req); err != nil {
        log.Println(err)
    }

    s := service.NewProductService(c.ctx, c.svcCtx)
    if err := s.CreateProduct(req); err != nil {
        ctx.JSON(http.StatusInternalServerError, gate.ErrorInfo{
            Code:   http.StatusInternalServerError,
            Reason: err.Error(),
            Msg:    "",
        })
        return
    }

    ctx.JSON(http.StatusOK, gate.CreateProductResp{})
}

func (c *ProductController) QueryPopularProduct(ctx *gin.Context) {
    req := &gate.QueryPopularProductReq{}
    if err := ctx.ShouldBindJSON(req); err != nil {
        log.Println(err)
    }

    s := service.NewProductService(c.ctx, c.svcCtx)
    resp, err := s.QueryPopularProduct(req)
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
