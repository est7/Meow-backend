package main

import (
	"Meow-backend/initialize"
	"Meow-backend/pkg/log"
	"context"
	_ "github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag"
	_ "github.com/thanhpk/randstr"
)

func main() {
	ctx := context.Background()

	// Load config
	conf := initialize.LoadConfig(initialize.ConfigPath)
	initialize.LoadLoggerConfig(initialize.ConfigPath)

	log.Debugf("success initialize logg")

	// Initialize database
	gormDB, db, err := initialize.InitDB(conf.PGConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	initialize.Instance.Db = db
	initialize.Instance.GormDb = gormDB

	defer func() {
		if err := initialize.CloseDB(db); err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}()

	// Initialize Redis
	redisClient, err := initialize.InitRedis(conf.RedisConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	initialize.Instance.RedisClient = redisClient

	defer func() {
		if err := initialize.CloseRedis(redisClient); err != nil {
			log.Fatalf("Error closing Redis: %v", err)
		}
	}()

	// Initialize route
	r, err := initialize.InitRoute(ctx)
	if err != nil {
		log.Fatalf("Error initializing route: %v", err)
	}

	// Start the server
	log.Infof("Server started on port %s", conf.Port)
	if err := r.Run(conf.Port); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
