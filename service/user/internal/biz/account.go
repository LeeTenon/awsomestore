package biz

import (
    "awsomestore/api/user"
    "awsomestore/service/user/internal/model"
    "context"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
)

var (
    ErrUserNotFound = errors.NotFound(user.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type AccountRepo interface {
    First(context.Context, string) (*model.Account, error)
    Insert(ctx context.Context, account *model.Account) error
}

type AccountUsecase struct {
    repo AccountRepo
    log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
    return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUsecase) QueryUser(ctx context.Context, email string) (*model.Account, error) {
    uc.log.WithContext(ctx).Infof("Query user with email: %s", email)
    result, err := uc.repo.First(ctx, email)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (uc *AccountUsecase) CreateUser(ctx context.Context, account *model.Account) error {
    uc.log.WithContext(ctx).Infof("Insert user %v", account)
    err := uc.repo.Insert(ctx, account)

    return err
}
