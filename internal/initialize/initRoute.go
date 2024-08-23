package initialize

import (
	"Meow-backend/internal/modules"
	"Meow-backend/pkg/app"
	"Meow-backend/pkg/errcode"
	"Meow-backend/pkg/middlewares"
	"Meow-backend/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(ctx context.Context, appCtx *AppInstance) (*gin.Engine, error) {
	// 设置 Gin 路由
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	// 遍历并将每个中间件添加到 Gin 服务器
	for _, middleware := range middlewares.Middleware {
		server.Use(middleware)
	}

	// load web router
	LoadWebRouter(server)

	server.NoRoute(func(c *gin.Context) {
		app.ErrorResponse(c, errcode.ErrApiNotFound)
	})

	server.NoMethod(func(c *gin.Context) {
		app.ErrorResponse(c, errcode.ErrMethodNotAllowed)
	})

	// HealthCheck 健康检查路由
	server.GET("/health", healthCheck)
	// hostnameHealthCheck 主机名健康检查路由
	server.GET("/hostname", app.HostnameHealthCheck)

	allModules := modules.InitModules(appCtx)
	for _, module := range allModules {
		module.RegisterRoutes(server, middlewares.CreateAuthMiddleware)
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
	if err := AppCtxInstance.Db.Ping(); err != nil {
		status = "DOWN"
		details["database"] = "Database connection failed: " + err.Error()
	} else {
		details["database"] = "Connected"
	}

	// 检查 Redis 连接
	if _, err := AppCtxInstance.RedisClient.Ping(c).Result(); err != nil {
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
		app.ErrorResponse(c, errcode.NewCustomError(http.StatusServiceUnavailable, "Service is unhealthy"))
	}
}
