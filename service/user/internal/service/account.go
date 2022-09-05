package service

import (
    pb "awsomestore/api/user"
    "awsomestore/common/id"
    "awsomestore/service/user/internal/biz"
    "awsomestore/service/user/internal/model"
    "context"
    "fmt"
    "strconv"
)

type AccountService struct {
    pb.UnimplementedAccountServer

    uc *biz.AccountUsecase
}

func NewAccountService(uc *biz.AccountUsecase) *AccountService {
    fmt.Println("init")
    return &AccountService{
        uc: uc,
    }
}

func (s *AccountService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    uid, err := id.Gen()
    if err != nil {
        return nil, err
    }

    if err = s.uc.CreateUser(ctx, &model.Account{
        Uid:       strconv.FormatInt(uid, 10),
        Name:      req.Name,
        Email:     req.Email,
        Password:  req.Password,
        AvatarUrl: req.Avatar,
    }); err != nil {
        return nil, err
    }

    return &pb.CreateUserResp{}, nil
}
func (s *AccountService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
    return &pb.UpdateUserResp{}, nil
}
func (s *AccountService) QueryUser(ctx context.Context, req *pb.QueryUserReq) (*pb.QueryUserResp, error) {
    result, err := s.uc.QueryUser(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    return &pb.QueryUserResp{
        Name:     result.Name,
        Email:    result.Email,
        Password: result.Password,
        Avatar:   result.AvatarUrl,
    }, nil
}
func (s *AccountService) ListUserProfile(ctx context.Context, req *pb.ListUserProfileReq) (*pb.ListUserProfileResp, error) {
    return &pb.ListUserProfileResp{}, nil
}
