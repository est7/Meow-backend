package initialize

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(ctx context.Context) (*gin.Engine, error) {
	// 设置 Gin 路由
	r := gin.Default()

	// 定义路由
	r.GET("/", helloWorld)
	r.GET("/health", healthCheck)

	return r, nil
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

// healthCheck 处理健康检查请求
func healthCheck(c *gin.Context) {
	// 检查数据库连接
	//err := AppInstance{hhh}.Ping()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection failed"})
	//	return
	//}
	//
	//// 检查 Redis 连接
	//_, err = rdb.Ping(c).Result()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Redis connection failed"})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is healthy"})
}
