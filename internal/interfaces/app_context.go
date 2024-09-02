package interfaces

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDB() *sql.DB
	GetGormDB() *gorm.DB
	GetRedisClient() *redis.Client
	// 添加其他需要的方法
	//GetConfig() AppEnvConfig
	//GetLogger() *log.Logger
}
