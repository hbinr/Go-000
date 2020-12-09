package service

import (
	"hb.study/Week05/code/work/internal/user/model"
)

func (u *userService) GetUser(userID int64) (*model.User, error) {
	return u.dao.SelectByUserID(userID)
}
