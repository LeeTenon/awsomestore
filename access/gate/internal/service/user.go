package service

import (
    "awsomestore/access/gate/internal/svc"
    "awsomestore/api/gate"
    pb "awsomestore/api/user"
    "context"
    "log"
)

type UserService struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUserService(ctx context.Context, svcCtx *svc.ServiceContext) *UserService {
    return &UserService{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (s *UserService) CreateUser(req *gate.CreateUserReq) error {
    _, err := s.svcCtx.User.Account.CreateUser(s.ctx, &pb.CreateUserReq{
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password,
        Avatar:   req.Avatar,
    })
    if err != nil {
        log.Println(err)
    }

    return err
}

func (s *UserService) QueryCart(req *gate.QueryCartReq) (*gate.QueryCartResp, error) {
    resp, err := s.svcCtx.User.Cart.QueryCart(s.ctx, &pb.QueryCartReq{
        Uid: req.Uid,
    })
    if err != nil {
        log.Println(err)
    }

    items := make([]*gate.CartItem, 0, 5)
    for _, item := range resp.CartItems {
        items = append(items, &gate.CartItem{
            Pid:      item.Pid,
            Title:    item.Title,
            Type:     item.Type,
            Count:    item.Count,
            Price:    item.Price,
            CoverUrl: item.CoverUrl,
        })
    }

    return &gate.QueryCartResp{
        Items: items,
    }, nil
}
