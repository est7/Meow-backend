package interfaces

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//type Module interface {
//	Name() string
//	Init(appCtx AppContext)
//	RegisterRoutes(r *gin.Engine, authMiddleware func(auth.PermissionLevel) gin.HandlerFunc)
//}

type AppContext interface {
	GetDB() *sql.DB
	GetGormDB() *gorm.DB
	GetRedisClient() *redis.Client
	// 添加其他需要的方法
	//GetConfig() AppEnvConfig
	//GetLogger() *log.Logger
}
