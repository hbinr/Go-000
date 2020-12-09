package dao

import (
	"database/sql"
	"fmt"

	"hb.study/Week05/code/work/internal/user/model"

	"github.com/pkg/errors"
)

const (
	_selectByUserIDSQL = `select user_id,user_name from user where user_id = ?`
)

func (u *userDao) SelectByUserID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	fmt.Println("boolssssssssssssssss", user == nil)
	if err = u.db.Get(user, _selectByUserIDSQL, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
			user = nil
			return
		}
		err = errors.Wrapf(err, "dao select user by user_id:%d", userID)
	}
	return
}
