package entity

// User user 领域实体
type User struct {
	UserID   int64
	Passport string
	Password string
	Email    string
	Nickname string
}

// 贫血模型设计：增加一些简单的逻辑判断，不做持久化操作

// CheckPassport 查询账号唯一性
func (do *User) CheckPassport() bool {
	// 查表判断，账号是否存在
	return true
}

// CheckNickName 查询昵称唯一性
func (do *User) CheckNickName() bool {
	// 查表判断，昵称是否存在
	return true
}

// CheckEmail 查询邮箱唯一性
func (do *User) CheckEmail() bool {
	// 查表判断，邮箱是否存在
	return true
}
