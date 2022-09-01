package data

import (
    model "awsomestore/models/user"
    "awsomestore/services/user/internal/biz"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    if err := data.DB.AutoMigrate(&model.Account{}); err != nil {
        log.Errorf("auto migrate error: %s", err)
    }

    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (u *userRepo) First(ctx context.Context, email string) (*model.Account, error) {
    result := &model.Account{}
    if err := u.data.DB.Where("email", email).First(result).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, err
        }
        return nil, err
    }
    return result, nil
}

func (u *userRepo) Insert(ctx context.Context, account *model.Account) error {
    if err := u.data.DB.Create(account).Error; err != nil {
        return err
    }

    return nil
}

func (u *userRepo) Update(ctx context.Context, email string) error {
    // TODO
    return nil
}

func (u *userRepo) Find(ctx context.Context, email string) ([]*model.Account, error) {
    // TODO
    return nil, nil
}
