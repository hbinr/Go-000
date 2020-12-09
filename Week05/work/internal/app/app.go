package app

import (
	"fmt"

	"hb.study/Week05/code/work/conf"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"hb.study/Week05/code/work/internal/user/controller"
)

type WebApp struct {
	*gin.Engine
	Config   *conf.Config
	UserCtrl *controller.UserController
}

// InitEngine 初始化gin
func InitEngine(c *conf.Config) *gin.Engine {
	if c.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.Default()
	// 设置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置公共中间件
	//r.Use(
	//	middleware.GinLogger(),       // zap logger中间件
	//	middleware.GinRecovery(true), // zap recovery中间件
	//)
	return r
}

// Start the web app
func (e *WebApp) Start() {
	e.Run(fmt.Sprintf(":%d", e.Config.System.Port))
}
