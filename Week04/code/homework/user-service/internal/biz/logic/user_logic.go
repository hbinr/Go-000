package logic

import (
	"context"

	"user.service/pkg/tool/hash"

	"user.service/internal/pkg/codec"

	pb "user.service/api/user/v1"

	"user.service/pkg/tool/snowflake"

	"user.service/internal/biz/repository/po"

	"user.service/internal/biz/factory"

	"github.com/google/wire"
	"user.service/internal/biz/entity"
	"user.service/internal/biz/repository"
)

var UserLogicSet = wire.NewSet(wire.Struct(new(UserLogic), "*"))

type UserLogic struct {
	UserRepo repository.IUserRepo
}

// CreateUser 创建用户
func (u *UserLogic) CreateUser(_ context.Context, userDO *entity.User) (dto *pb.UserDTO, err error) {
	userPO := new(po.UserPO)
	userPO.UserID = snowflake.GenID()
	userPO = factory.NewUserPO(userDO) // do -> po
	// 1.校验账号
	if u.UserRepo.CheckPassport(userDO.Passport) {
		return nil, codec.CodeUserExist
	}
	// 2.密码加密
	userPO.Password = hash.MD5String(userDO.Password)

	// 3.数据落库
	err = u.UserRepo.CreateUser(userPO)

	dto = factory.NewUserDTO(userPO)
	return
}
