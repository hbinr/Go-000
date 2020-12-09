package dao

import (
	"github.com/jmoiron/sqlx"
	"hb.study/Week05/code/work/internal/user/model"
)

type (
	IUserDao interface {
		SelectByUserID(userID int64) (*model.User, error)
	}

	userDao struct {
		db *sqlx.DB
	}
)

func NewUserDao(db *sqlx.DB) IUserDao {
	return &userDao{
		db: db,
	}
}
