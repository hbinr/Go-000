package factory

import (
	pb "user.service/api/user/v1"
	"user.service/internal/biz/entity"
	"user.service/internal/biz/repository/po"
)

/*
	数据工厂，主要职责：
		创建对象，通过对象拼接返回需要的对象
*/
// NewUserDO 逻辑简单，单纯的创建对象，没有复杂的对象拼接
func NewUserDO(po *po.UserPO) *entity.User {
	return &entity.User{
		UserID:   po.UserID,
		Passport: po.Password,
		Password: po.Password,
		Email:    po.Email,
		Nickname: po.Nickname,
	}
}

// NewUserPO .
func NewUserPO(do *entity.User) *po.UserPO {
	return &po.UserPO{
		UserID:   do.UserID,
		Passport: do.Passport,
		Password: do.Password,
		Email:    do.Email,
		Nickname: do.Nickname,
	}
}

func NewUserDTO(po *po.UserPO) *pb.UserDTO {
	return &pb.UserDTO{
		UserID:     po.UserID,
		Passport:   po.Passport,
		Email:      po.Email,
		Nickname:   po.Nickname,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
