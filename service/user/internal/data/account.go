package data

import (
    "awsomestore/service/user/internal/biz"
    "awsomestore/service/user/internal/model"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type accountRepo struct {
    data *Data
    log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
    if err := data.DB.AutoMigrate(&model.Account{}); err != nil {
        log.Errorf("auto migrate error: %s", err)
    }

    return &accountRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (u *accountRepo) First(ctx context.Context, email string) (*model.Account, error) {
    result := &model.Account{}
    if err := u.data.DB.Where("email", email).First(result).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, biz.ErrUserNotFound
        }
        return nil, err
    }
    return result, nil
}

func (u *accountRepo) Insert(ctx context.Context, account *model.Account) error {
    if err := u.data.DB.Create(account).Error; err != nil {
        return err
    }

    return nil
}
