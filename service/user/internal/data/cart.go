package data

import (
    "awsomestore/service/user/internal/biz"
    "awsomestore/service/user/internal/model"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type cartRepo struct {
    data *Data
    log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
    if err := data.DB.AutoMigrate(&model.Cart{}); err != nil {
        log.Errorf("auto migrate error: %s", err)
    }

    return &cartRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (u *cartRepo) Update(ctx context.Context, cart *model.Cart) error {
    if err := u.data.DB.Model(&model.Cart{}).Where("uid", cart.Uid).Updates(map[string]interface{}{
        "pid_array":   cart.PidArray,
        "count_array": cart.CountArray,
    }).Error; err != nil {
        if err != gorm.ErrRecordNotFound {
            return err
        }
    }

    return nil
}

func (u *cartRepo) FindByUid(ctx context.Context, uid string) (*model.Cart, error) {
    result := &model.Cart{}

    if err := u.data.DB.Where("uid", uid).FirstOrCreate(result).Error; err != nil {
        return nil, err
    }

    return result, nil
}
