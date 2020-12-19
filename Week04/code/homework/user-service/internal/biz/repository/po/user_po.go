package po

type UserPO struct {
	Id         int64  `gorm:"primaryKey;autoIncrement"`    // 逐渐
	UserID     int64  `gorm:"not null;index:idx_user_id;"` // 用户ID
	Passport   string `gorm:"size:32;unique;"`             // 用户名
	Password   string `gorm:"size:512"`                    // 密码
	Email      string `gorm:"size:128;unique;"`            // 邮箱
	Nickname   string `gorm:"size:16;"`                    // 昵称
	CreateTime int64  `gorm:"autoCreateTime"`              // 创建时间, 使用时间戳秒数填充创建时间,方便前端自行设置时间格式
	UpdateTime int64  `gorm:"autoUpdateTime"`              // 更新时间
}

func (po *UserPO) TableName() string {
	return "user"
}
