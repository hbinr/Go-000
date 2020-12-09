package model

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"user_name"`
}
