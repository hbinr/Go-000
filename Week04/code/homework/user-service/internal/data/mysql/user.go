package mysql

import (
	"errors"

	"github.com/google/wire"
	"gorm.io/gorm"
	"user.service/internal/biz/repository"
	"user.service/internal/biz/repository/po"
)

/*
	数据持久化
*/
var (
	// 验证接口是否实现
	_ repository.IUserRepo = (*UserRepo)(nil)
	//UserRepoSet 使用wire 依赖注入
	UserRepoSet = wire.NewSet(
		wire.Struct(new(UserRepo), "*"),
		wire.Bind(new(repository.IUserRepo), new(*UserRepo)))
)

type UserRepo struct {
	DB *gorm.DB
}

// CreateUser .
func (u *UserRepo) CreateUser(po *po.UserPO) error {
	return u.DB.Create(po).Error
}

// CheckPassport 校验账户唯一性
func (u *UserRepo) CheckPassport(passport string) bool {
	existFlag := true
	var po po.UserPO
	if err := u.DB.Table(po.TableName()).Where("passport = ?", passport).First(&po).Error; err != nil {
		existFlag = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return existFlag
		}
		return existFlag
	}
	return existFlag
}
