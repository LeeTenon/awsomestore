package service

import (
    "awsomestore/api/product"
    "awsomestore/common/id"
    "awsomestore/service/product/internal/biz"
    "awsomestore/service/product/internal/model"
    "context"
    "strconv"
    "strings"
)

type ProductService struct {
    product.UnimplementedProductServer

    uc *biz.ProductUsecase
}

func NewProductService(uc *biz.ProductUsecase) *ProductService {
    return &ProductService{uc: uc}
}

func (s *ProductService) CreateProduct(ctx context.Context, in *product.CreateProductReq) (*product.CreateProductResp, error) {
    pid, err := id.Gen()
    if err != nil {
        return nil, err
    }
    keywords := strings.Join(in.Keyword, ",")
    pictures := strings.Join(in.Picture, ",")

    err = s.uc.CreateProduct(ctx, &model.Product{
        Pid:        strconv.FormatInt(pid, 10),
        Title:      in.Title,
        Desc:       in.Desc,
        Keyword:    keywords,
        Price:      in.Price,
        PictureUrl: pictures,
    })
    if err != nil {
        return nil, err
    }

    return &product.CreateProductResp{}, nil
}

func (s *ProductService) ListProduct(ctx context.Context, in *product.ListProductReq) (*product.ListProductResp, error) {
    result, err := s.uc.QueryProductList(ctx, in.Keyword)
    if err != nil {
        return nil, err
    }

    pbResult := make([]*product.ProductInfo, 0, 10)
    for _, item := range result {
        pbResult = append(pbResult, &product.ProductInfo{
            Pid:     item.Pid,
            Title:   item.Title,
            Desc:    item.Desc,
            Price:   item.Price,
            Picture: strings.Split(item.PictureUrl, ","),
        })
    }

    return &product.ListProductResp{
        ProductList: pbResult,
    }, nil
}
