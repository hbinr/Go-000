package controller

import (
	"github.com/gin-gonic/gin"
	"hb.study/Week05/code/work/internal/user/service"
)

type UserController struct {
	engine  *gin.Engine
	service service.IUserService
}

func NewUserController(e *gin.Engine, userService service.IUserService) *UserController {
	user := &UserController{
		engine:  e,
		service: userService,
	}
	g := e.Group("/api/user")
	{
		g.GET("/get", user.Get)
	}
	return user
}
