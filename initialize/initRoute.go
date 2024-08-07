package initialize

import (
	"Meow-backend/pkg/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(ctx context.Context) (*gin.Engine, error) {
	// 设置 Gin 路由
	server := gin.Default()
	// 遍历并将每个中间件添加到 Gin 服务器
	for _, middleware := range middlewares.Middleware {
		server.Use(middleware)
	}

	// load web router
	LoadWebRouter(server)

	server.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The route was not found")
	})

	server.NoMethod(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The route was not found")
	})

	// HealthCheck 健康检查路由
	server.GET("/health", healthCheck)

	return server, nil
}

func LoadWebRouter(server *gin.Engine) {

}

// healthCheck 处理健康检查请求
func healthCheck(c *gin.Context) {
	// 检查数据库连接

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
