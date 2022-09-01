package service

import (
    "awsomestore/access/gate/internal/biz"
    "context"

    pb "awsomestore/api/gate"
)

type UserService struct {
    pb.UnimplementedGateServer

    uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
    return &UserService{uc: uc}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
    // todo: 查询用户

    return &pb.LoginResp{}, nil
}
