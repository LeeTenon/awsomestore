package data

import (
    "awsomestore/service/product/internal/model"
    "context"
    "gorm.io/gorm"

    "awsomestore/service/product/internal/biz"
    "github.com/go-kratos/kratos/v2/log"
)

type ProductRepo struct {
    data *Data
    log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
    if err := data.DB.AutoMigrate(&model.Product{}); err != nil {
        log.Errorf("auto migrate error: %s", err)
    }

    return &ProductRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (r *ProductRepo) Save(ctx context.Context, p *model.Product) error {
    if err := r.data.DB.FirstOrCreate(p).Error; err != nil {
        r.log.Errorf("create product[%s] error: %s", p.Pid, err.Error())
        return err
    }

    return nil
}

func (r *ProductRepo) FindByKeyword(ctx context.Context, keyword string) ([]*model.Product, error) {
    result := make([]*model.Product, 10)
    if err := r.data.DB.Where("keyword LIKE ?", keyword).Find(&result).Error; err != nil {
        if err != gorm.ErrRecordNotFound {
            r.log.WithContext(ctx).Errorf("query product by keyword[%s] error: %s", keyword, err.Error())
            return nil, err
        }
        return nil, nil
    }

    return result, nil
}

//func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
//	return g, nil
//}
//
//func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
//	return nil, nil
//}
