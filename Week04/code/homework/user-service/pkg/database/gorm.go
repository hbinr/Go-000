package database

import (
	"database/sql"
	"os"

	"user.service/internal/biz/repository/po"

	"github.com/apache/dubbo-go/common/logger"

	"user.service/pkg/conf"

	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

var (
	sqlDB *sql.DB
)

// InitMySQL 初始化MySQL
func InitMySQL(cfg *conf.MyConfig) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       cfg.DSN, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	gormConfig := config(cfg.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		logger.Error("opens database failed", err)
		return
	}

	if sqlDB, err = db.DB(); err != nil {
		logger.Error("db.db() failed", err)
		return
	}
	gormDBTables(db)

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	return
}

// gormDBTables 注册数据库表专用
func gormDBTables(db *gorm.DB) {
	err := db.AutoMigrate(&po.UserPO{})
	if err != nil {
		logger.Error("register table failed", err)
		os.Exit(0)
	}
	logger.Info("register table success")
}

// config 根据配置决定是否开启日志
func config(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   gormLog.Default.LogMode(gormLog.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   gormLog.Default.LogMode(gormLog.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
