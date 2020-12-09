package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"hb.study/Week05/code/work/internal/user"
	"hb.study/Week05/code/work/pkg/database"
)

func main() {
	app, _ := initApp()
	app.Run() // 默认8080端口
}

// initApp 初始化
func initApp() (e *gin.Engine, err error) {
	var db *sqlx.DB
	if db, err = database.InitMysql(); err != nil {
		return nil, err
	}
	r := gin.Default()
	user.Build(r, db)
	return
}
