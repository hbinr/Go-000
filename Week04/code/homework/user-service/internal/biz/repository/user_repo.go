package repository

import "user.service/internal/biz/repository/po"

/*
	User接口定义
*/

type IUserRepo interface {
	CreateUser(po *po.UserPO) error
	CheckPassport(passport string) bool
}
