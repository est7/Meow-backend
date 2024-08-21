package main

import (
	"Meow-backend/initialize"
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/log"
	"context"
	"database/sql"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag"
	_ "github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	conf := initConfig()
	initModules()

	db, gormDB := initDatabase(ctx, conf)
	gormDB.AutoMigrate()

	defer initialize.CloseDB(db)

	redisClient := initRedis(ctx, conf)
	defer initialize.CloseRedis(redisClient)

	r := initRouter(ctx)

	startServer(r, conf.Port)
}

func initConfig() *initialize.AppEnvConfig {
	conf := initialize.LoadConfig(initialize.ConfigPath)
	initialize.LoadLoggerConfig(initialize.ConfigPath, conf.Mode)
	log.Debugf("success initialize log")
	return &conf
}

func initModules() {
	for _, m := range modules.Modules {
		m.Init()
	}
}

func initDatabase(ctx context.Context, conf *initialize.AppEnvConfig) (*sql.DB, *gorm.DB) {
	gormDB, db, err := initialize.InitDB(conf.PGConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	initialize.Instance.Db = db
	initialize.Instance.GormDb = gormDB
	return db, gormDB
}

func initRedis(ctx context.Context, conf *initialize.AppEnvConfig) *redis.Client {
	redisClient, err := initialize.InitRedis(conf.RedisConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	initialize.Instance.RedisClient = redisClient
	return redisClient
}

func initRouter(ctx context.Context) *gin.Engine {
	r, err := initialize.InitRoute(ctx)
	if err != nil {
		log.Fatalf("Error initializing route: %v", err)
	}
	return r
}

func startServer(r *gin.Engine, port string) {
	log.Infof("Server started on port %s", port)
	if err := r.Run(port); err != nil {
		panic(err)
		log.Fatalf("Error running server: %v", err)
	}
}
