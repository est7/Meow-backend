package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

var path = "config/"

func main() {
	//hello,world
	fmt.Println("hello,world")
	//ctx := context.Background()
	/*
		// Load config
		conf := initialize.LoadConfig(path)

		// Initialize database
		db, err := initialize.InitDB(conf.PGConfig, ctx)
		if err != nil {
			log.Fatalf("Error initializing database: %v", err)
		}
		initialize.Instance.Db = db
		defer initialize.CloseDB(db)

		// Initialize Redis
		redis, err := initialize.InitRedis(conf.RedisConfig, ctx)
		if err != nil {
			log.Fatalf("Error initializing Redis: %v", err)
		}
		initialize.Instance.RedisClient = redis
		defer initialize.CloseRedis(redis)

		// Initialize route
		r, err := initialize.InitRoute(ctx)
		if err != nil {
			log.Fatalf("Error initializing route: %v", err)
		}

		// Start the server
		log.Printf("Server started on port %s", conf.Port)
		if err := r.Run(conf.Port); err != nil {
			log.Fatalf("Error running server: %v", err)
		}*/
}
