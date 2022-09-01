package service

import (
    pb "awsomestore/api/user"
    "awsomestore/common/id"
    "awsomestore/models/user"
    "awsomestore/services/user/internal/biz"
    "context"
    "strconv"
)

type UserService struct {
    pb.UnimplementedUserServer

    uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
    return &UserService{
        uc: uc,
    }
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    uid, err := id.Gen()
    if err != nil {
        return nil, err
    }

    if err = s.uc.CreateUser(ctx, &user.Account{
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
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
    return &pb.UpdateUserResp{}, nil
}
func (s *UserService) QueryUser(ctx context.Context, req *pb.QueryUserReq) (*pb.QueryUserResp, error) {
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
func (s *UserService) ListUserProfile(ctx context.Context, req *pb.ListUserProfileReq) (*pb.ListUserProfileResp, error) {
    return &pb.ListUserProfileResp{}, nil
}
