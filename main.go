package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	db  *sql.DB
	rdb *redis.Client
	err error
)

func main() {
	// 初始化数据库连接
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 初始化 Redis 客户端
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	// 设置 Gin 路由
	r := gin.Default()

	// 定义路由
	r.GET("/", helloWorld)
	r.GET("/health", healthCheck)

	// 启动服务器
	r.Run(":8080")
}

// helloWorld 处理根路径请求
func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// healthCheck 处理健康检查请求
func healthCheck(c *gin.Context) {
	// 检查数据库连接
	err := db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection failed"})
		return
	}

	// 检查 Redis 连接
	_, err = rdb.Ping(c).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Redis connection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is healthy"})
}
