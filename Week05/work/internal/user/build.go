package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"hb.study/Week05/code/work/internal/user/controller"
	"hb.study/Week05/code/work/internal/user/service"
)

func Build(engine *gin.Engine, db *sqlx.DB) *controller.UserController {
	us := service.NewUserService(db)
	return controller.NewUserController(engine, us)
}
