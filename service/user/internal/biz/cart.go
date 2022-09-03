package biz

import (
    pb "awsomestore/api/user"
    "context"
    "github.com/go-kratos/kratos/v2/errors"
    "strconv"
    "strings"

    "awsomestore/service/user/internal/model"

    "github.com/go-kratos/kratos/v2/log"
)

type CartRepo interface {
    Update(context.Context, *model.Cart) error
    FindByUid(ctx context.Context, uid string) (*model.Cart, error)
}

type CartUsecase struct {
    repo CartRepo
    log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
    return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) UpdateCart(ctx context.Context, uid string, items []*pb.CartItem) error {
    uc.log.WithContext(ctx).Infof("Update cart with items: %s", items)
    uc.log.Infof("Update cart with items: %s", items)

    pidArray := make([]string, 0, 5)
    countArray := make([]string, 0, 5)
    for _, item := range items {
        pidArray = append(pidArray, item.Pid)
        countArray = append(countArray, strconv.Itoa(int(item.Count)))
    }
    err := uc.repo.Update(ctx, &model.Cart{
        Uid:        uid,
        PidArray:   strings.Join(pidArray, ","),
        CountArray: strings.Join(countArray, ","),
    })
    if err != nil {
        return err
    }

    return nil
}

func (uc *CartUsecase) QueryCart(ctx context.Context, uid string) (map[string]int32, error) {
    cart, err := uc.repo.FindByUid(ctx, uid)
    if err != nil {
        return nil, err
    }

    items := make(map[string]int32)
    pids := strings.Split(cart.PidArray, ",")
    counts := strings.Split(cart.CountArray, ",")
    if pids[0] == "" {
        return nil, nil
    }
    if len(pids) != len(counts) {
        uc.log.WithContext(ctx).Errorf("uid[%s] 条目损毁!", uid)
        // TODO:可承受错误，提供自修复
        return nil, errors.InternalServer("条目损毁", "条目损毁")
    }
    for i := 0; i < len(pids); i++ {
        count, _ := strconv.Atoi(counts[i])
        items[pids[i]] = int32(count)
    }

    return items, nil
}
