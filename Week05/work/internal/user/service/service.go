package service

import (
	"github.com/jmoiron/sqlx"
	"hb.study/Week05/code/work/internal/user/dao"
	"hb.study/Week05/code/work/internal/user/model"
)

type (
	IUserService interface {
		GetUser(userID int64) (*model.User, error)
	}

	userService struct {
		dao dao.IUserDao
	}
)

func NewUserService(db *sqlx.DB) IUserService {
	return &userService{
		dao: dao.NewUserDao(db),
	}
}
