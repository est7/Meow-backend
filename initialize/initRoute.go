package initialize

import (
	"Meow-backend/internal/module"
	"Meow-backend/pkg/app"
	"Meow-backend/pkg/errcode"
	"Meow-backend/pkg/middlewares"
	"Meow-backend/pkg/utils"
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
		app.Error(c, errcode.ErrApiNotFound)
	})

	server.NoMethod(func(c *gin.Context) {
		app.Error(c, errcode.ErrMethodNotAllowed)
	})

	// HealthCheck 健康检查路由
	server.GET("/health", healthCheck)
	// hostnameHealthCheck 主机名健康检查路由
	server.GET("/hostname", app.HostnameHealthCheck)

	apiV1Pub := server.Group("/v1")

	apiV1Pri := server.Group("/v1")
	apiV1Pri.Use(middlewares.JWT())

	for _, m := range module.Modules {
		m.InitRouter(apiV1Pub, apiV1Pri)
	}

	return server, nil
}

func LoadWebRouter(server *gin.Engine) {
	//todo
}

// healthCheck 处理健康检查请求
func healthCheck(c *gin.Context) {
	status := "UP"
	details := make(map[string]string)

	// 检查数据库连接
	if err := Instance.Db.Ping(); err != nil {
		status = "DOWN"
		details["database"] = "Database connection failed: " + err.Error()
	} else {
		details["database"] = "Connected"
	}

	// 检查 Redis 连接
	if _, err := Instance.RedisClient.Ping(c).Result(); err != nil {
		status = "DOWN"
		details["redis"] = "Redis connection failed: " + err.Error()
	} else {
		details["redis"] = "Connected"
	}

	healthData := gin.H{
		"status":   status,
		"details":  details,
		"hostname": utils.GetHostname(),
	}

	if status == "UP" {
		app.SuccessResponse(c, healthData)
	} else {
		app.Error(c, errcode.NewCustomError(http.StatusServiceUnavailable, "Service is unhealthy"))
	}
}
