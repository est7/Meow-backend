package middlewares

import (
	"Meow-backend/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAuthMiddleware(level auth.PermissionLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 根据不同的权限级别实现相应的逻辑
		switch level {
		case auth.Public:
			// 无需验证，直接通过
			c.Next()
		case auth.Authenticated:
			// 验证用户是否登录
			if !isAuthenticated(c) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		case auth.Admin:
			// 验证用户是否为管理员
			if !isAdmin(c) {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
		c.Next()
	}
}

func isAuthenticated(c *gin.Context) bool {
	// 实现用户认证逻辑
}

func isAdmin(c *gin.Context) bool {
	// 实现管理员验证逻辑
}
