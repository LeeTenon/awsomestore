package data

import (
    "context"

    "awsomestore/access/gate/internal/biz"
    "github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (u *userRepo) First(ctx context.Context, email string) (*biz.User, error) {
    // TODO
    return nil, nil
}

func (u *userRepo) Insert(ctx context.Context, email string)  error {
    // TODO
    return nil
}

func (u *userRepo) Update(ctx context.Context, email string) error {
    // TODO
    return nil
}


func (u *userRepo) Find(ctx context.Context, email string) ([]*biz.User, error) {
    // TODO
    return nil, nil
}