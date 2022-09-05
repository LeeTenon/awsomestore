package biz

import (
    "context"

    "github.com/go-kratos/kratos/v2/log"

    "awsomestore/service/product/internal/model"
)

type ProductRepo interface {
    Save(context.Context, *model.Product) error
    FindByKeyword(context.Context, string) ([]*model.Product, error)
}

type ProductUsecase struct {
    repo ProductRepo
    log  *log.Helper
}

func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
    return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, g *model.Product) error {
    return uc.repo.Save(ctx, g)
}

func (uc *ProductUsecase) QueryProductList(ctx context.Context, keyword string) ([]*model.Product, error) {
    return uc.repo.FindByKeyword(ctx, keyword)
}
