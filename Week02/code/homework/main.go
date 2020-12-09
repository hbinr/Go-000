package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const _selectUserByUserIDSQL = `select user_isd,user_name from user where user_id = ?`

var (
	db                *sqlx.DB
	ErrRecordNotFound = errors.New("dao: record not found")
)

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"user_name"`
}

func Dao(userID int64) (user *User, err error) {
	user = new(User)

	if err = db.Get(user, _selectUserByUserIDSQL, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Wrapf(ErrRecordNotFound, "by userID:(%d)", userID)
		} else {
			err = errors.Wrapf(err, "by userID:(%d)", userID)
		}
		return
	}
	return
}

func Biz(userID int64) (user *User, err error) {
	return Dao(userID)
}

func Http() {
	res, err := Biz(1)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			log.Println("no match data:", res)
			return
		}
		log.Printf("get user err:%+v", err)
		return
	}
	fmt.Println("res:", res)
}

func main() {
	if err := initDB(); err != nil {
		log.Printf("init db failed:%+v \r\n", err)
		return
	}
	Http()
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		return
	}
	return
}
