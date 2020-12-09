package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// InitMysql 初始化MySQL连接
func InitMysql() (db *sqlx.DB, err error) {
	//db = &sqlx.DB{}
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
	// sqlx.Connect() 底层做了Open和Ping
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		fmt.Println("sqlx.Connect failed, err:", err)
		return
	}
	// 设置MySQL相关配置
	db.SetConnMaxLifetime(20)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	fmt.Println("mysql connect success......")
	return
}
