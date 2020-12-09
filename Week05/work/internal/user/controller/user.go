package controller

import (
	"hb.study/Week05/code/work/internal/user/model"
	"hb.study/Week05/code/work/pkg/ginx"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Get(c *gin.Context) {
	var (
		userID int64
		err    error
		user   *model.User
	)
	if userID, err = ginx.QueryInt("userID", c); err != nil {
		ginx.FailWithMessage("invalid param", c)
		return
	}
	if user, err = u.service.GetUser(userID); err != nil {
		ginx.Fail(c)
		return
	}
	ginx.OkWithData(user, c)
}
