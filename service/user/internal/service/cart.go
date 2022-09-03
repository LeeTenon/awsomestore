package service

import (
    pb "awsomestore/api/user"
    "awsomestore/service/user/internal/biz"
    "context"
)

type CartService struct {
    pb.UnimplementedCartServer

    uc *biz.CartUsecase
}

func NewCartService(uc *biz.CartUsecase) *CartService {
    return &CartService{
        uc: uc,
    }
}

func (s *CartService) UpdateCart(ctx context.Context, req *pb.UpdateCartReq) (*pb.UpdateCartResp, error) {
    if err := s.uc.UpdateCart(ctx, req.Uid, req.CartItems); err != nil {
        return nil, err
    }

    return &pb.UpdateCartResp{}, nil
}
func (s *CartService) QueryCart(ctx context.Context, req *pb.QueryCartReq) (*pb.QueryCartResp, error) {
    result, err := s.uc.QueryCart(ctx, req.Uid)
    if err != nil {
        return nil, err
    }

    item := make([]*pb.CartItem, 0, 5)
    for k, v := range result {
        item = append(item, &pb.CartItem{
            Pid:   k,
            Count: v,
        })
    }

    return &pb.QueryCartResp{
        CartItems: item,
    }, nil
}
