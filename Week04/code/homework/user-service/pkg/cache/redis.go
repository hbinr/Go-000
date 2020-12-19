package cache

import (
	"fmt"

	"github.com/apache/dubbo-go/common/logger"

	"user.service/pkg/conf"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init 初始化redis连接
func InitRedis(cfg *conf.MyConfig) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default db
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		logger.Error("redis ping failed", err)
		return nil, err
	}
	return rdb, nil
}

// Close 关闭redis clent连接资源
func Close() {
	_ = rdb.Close()
}
