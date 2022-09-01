package biz

import (
    "awsomestore/api/user"
    model "awsomestore/models/user"
    "context"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
)

var (
    ErrUserNotFound = errors.NotFound(user.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserRepo interface {
    First(context.Context, string) (*model.Account, error)
    Insert(ctx context.Context, account *model.Account) error
}

type UserUsecase struct {
    repo UserRepo
    log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) QueryUser(ctx context.Context, email string) (*model.Account, error) {
    uc.log.WithContext(ctx).Infof("Query user with email: %s", email)
    result, err := uc.repo.First(ctx, email)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, account *model.Account) error {
    uc.log.WithContext(ctx).Infof("Insert user %v", account)
    err := uc.repo.Insert(ctx, account)

    return err
}
