package service

import (
	"context"

	"github.com/apache/dubbo-go/common/logger"

	pb "user.service/api/user/v1"
	"user.service/internal/biz/logic"
	"user.service/internal/service/assemble"
)

var (
	// 验证接口是否实现
	_ pb.UserServiceServer = (*UserServiceProvider)(nil)
)

type UserServiceProvider struct {
	*pb.UserServiceProviderBase
	userLgc *logic.UserLogic
}

func NewUserServiceProvider(userLgc *logic.UserLogic) *UserServiceProvider {
	return &UserServiceProvider{
		UserServiceProviderBase: &pb.UserServiceProviderBase{},
		userLgc:                 userLgc,
	}
}

func (u *UserServiceProvider) Register(ctx context.Context, req *pb.RegisterReq) (res *pb.RegisterRes, err error) {
	var (
		dto *pb.UserDTO
	)
	res = new(pb.RegisterRes)
	if dto, err = u.userLgc.CreateUser(ctx, assemble.Dto2Do(req)); err != nil {
		logger.Error("service: SignUp failed,err:", err)
		return
	}
	res.UserDTO = dto
	return
}

// Reference 必须实现RPCService接口，dubbo-go底层会调用
func (u *UserServiceProvider) Reference() string {
	return "UserServiceProvider"
}
