package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	Host         string        `mapstructure:"Host"`
	Port         int           `mapstructure:"Port"`
	Password     string        `mapstructure:"Password"`
	DB           int           `mapstructure:"DB"`
	MinIdleConn  int           `mapstructure:"MinIdleConn"`
	DialTimeout  time.Duration `mapstructure:"DialTimeout"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
	PoolSize     int           `mapstructure:"PoolSize"`
	PoolTimeout  time.Duration `mapstructure:"PoolTimeout"`
	EnableTrace  bool          `mapstructure:"EnableTrace"`
}

func InitRedis(redisConfig RedisConfig, ctx context.Context) (*redis.Client, error) {
	// 初始化 Redis 客户端
	redisHost := redisConfig.Host
	redisPort := redisConfig.Port

	RedisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		MinIdleConns: redisConfig.MinIdleConn,
		DialTimeout:  redisConfig.DialTimeout,
		ReadTimeout:  redisConfig.ReadTimeout,
		WriteTimeout: redisConfig.WriteTimeout,
		PoolSize:     redisConfig.PoolSize,
		PoolTimeout:  redisConfig.PoolTimeout,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err == nil {
		fmt.Println("Connected to Redis!")
	}
	return RedisClient, err

}

func CloseRedis(client *redis.Client) error {
	return client.Close()
}
