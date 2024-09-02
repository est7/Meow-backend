package main

import (
	"Meow-backend/internal/initialize"
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/modules/v1/card"
	"Meow-backend/internal/modules/v1/feed"
	"Meow-backend/internal/modules/v1/im"
	"Meow-backend/internal/modules/v1/user"
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

	db, gormDB := initDatabase(ctx, conf)
	gormDB.AutoMigrate()

	defer initialize.CloseDB(db)

	redisClient := initRedis(ctx, conf)
	defer initialize.CloseRedis(redisClient)

	var appCtxInstance = initialize.NewAppInstance(
		initialize.WithDB(db),
		initialize.WithGormDB(gormDB),
		initialize.WithRedisClient(redisClient),
	)
	r := initRouter(ctx, appCtxInstance)

	startServer(r, conf.Port)
}

func initConfig() *initialize.AppEnvConfig {
	conf := initialize.LoadConfig(initialize.ConfigPath)
	initialize.LoadLoggerConfig(initialize.ConfigPath, conf.Mode)
	log.Debugf("success initialize log")
	return &conf
}

func registerModule() {
	interfaces.RegisterModuleFactory(user.NewUserModule)
	interfaces.RegisterModuleFactory(feed.NewFeedModule)
	interfaces.RegisterModuleFactory(card.NewCardModule)
	interfaces.RegisterModuleFactory(im.NewIMModule)
}

func initDatabase(ctx context.Context, conf *initialize.AppEnvConfig) (*sql.DB, *gorm.DB) {
	gormDB, db, err := initialize.InitDB(conf.PGConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	return db, gormDB
}

func initRedis(ctx context.Context, conf *initialize.AppEnvConfig) *redis.Client {
	redisClient, err := initialize.InitRedis(conf.RedisConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	return redisClient
}

func initRouter(ctx context.Context, appCtx interfaces.AppContext) *gin.Engine {
	r, err := initialize.InitRoute(ctx, appCtx)
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

func init() {
	registerModule()
}
