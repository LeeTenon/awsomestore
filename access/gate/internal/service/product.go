package service

import (
    "awsomestore/access/gate/internal/svc"
    "awsomestore/api/gate"
    pb "awsomestore/api/product"
    "context"
    "log"
)

type ProductService struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewProductService(ctx context.Context, svcCtx *svc.ServiceContext) *ProductService {
    return &ProductService{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (s *ProductService) CreateProduct(req *gate.CreateProductReq) error {
    _, err := s.svcCtx.Product.CreateProduct(s.ctx, &pb.CreateProductReq{
        Title:    req.Title,
        Desc:     req.Desc,
        Category: req.Category,
        Price:    req.Price,
        Cover:    req.CoverUrl,
    })
    if err != nil {
        log.Println(err)
    }

    return err
}

func (s *ProductService) QueryPopularProduct(req *gate.QueryPopularProductReq) (*gate.QueryPopularProductResp, error) {
    resp, err := s.svcCtx.Product.ListProduct(s.ctx, &pb.ListProductReq{Keyword: "测试"})
    if err != nil {
        log.Println(err)
    }

    items := make([]*gate.ProductItem, 0, 5)
    for _, item := range resp.ProductList {
        items = append(items, &gate.ProductItem{
            Pid:      item.Pid,
            Title:    item.Title,
            Category: item.Category,
            Price:    item.Price,
            CoverUrl: item.Cover,
        })
    }

    return &gate.QueryPopularProductResp{
        ProductList: items,
    }, nil
}
