//+build wireinject

package main

import (
	"github.com/google/wire"
	"user.service/internal/biz/logic"
	"user.service/internal/data/mysql"
	"user.service/pkg/conf"
	"user.service/pkg/database"
)

// initLogic 注入函数，自定义的函数直接注入就行，不需要使用wire set
func initLogic() (*logic.UserLogic, error) {
	// 逻辑顺序入参，未用到的依赖不需要注入
	wire.Build(
		conf.Init,          // 初始化自定义conf，自定义
		database.InitMySQL, // 获取gorm.DB，自定义
		mysql.UserRepoSet,  // repository provider，wire生成
		logic.UserLogicSet, // logic UseUserLogicr，wire生成
	)
	// 返回值不用管。直接返回nil就行
	return nil, nil
}
