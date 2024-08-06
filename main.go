package main

import (
	"Meow-backend/initialize"
	"context"
	"database/sql"
	_ "github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag"
	_ "github.com/thanhpk/randstr"
	"log"
)

func main() {
	ctx := context.Background()

	// Load config
	conf := initialize.LoadConfig(initialize.ConfigPath)
	Logg = initialize.LoadLoggerConfig(initialize.ConfigPath)
	Logg.Debugf("success initialize config: %+v", conf)

	// Initialize database
	gormDB, db, err := initialize.InitDB(conf.PGConfig, ctx)
	if err != nil {
		Logg.Fatalf("Error initializing database: %v", err)
	}
	initialize.Instance.Db = db
	initialize.Instance.GormDb = gormDB

	defer func(client *sql.DB) {
		err := initialize.CloseDB(client)
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(db)

	// Initialize Redis
	redisClient, err := initialize.InitRedis(conf.RedisConfig, ctx)
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	initialize.Instance.RedisClient = redisClient

	defer func(db *redis.Client) {
		err := initialize.CloseRedis(redisClient)
		if err != nil {
			log.Fatalf("Error closing Redis: %v", err)
		}
	}(redisClient)

	// Initialize route
	r, err := initialize.InitRoute(ctx)
	if err != nil {
		log.Fatalf("Error initializing route: %v", err)
	}

	// Start the server
	log.Printf("Server started on port %s", conf.Port)
	if err := r.Run(conf.Port); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
