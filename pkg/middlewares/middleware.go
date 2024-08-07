package middlewares

import (
	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
	"net/http"
	"time"
)

var Middleware = defaultMiddlewares()

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"recovery":   gin.Recovery(),
		"secure":     Secure,
		"options":    Options,
		"nocache":    NoCache,
		"logger":     Logging(),
		"cors":       Cors(),
		"request_id": RequestID(),
	}
}

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		//OPTIONS请求通常被称为预检请求，它是在实际请求（如GET、POST等）之前发送的，以确定服务器是否允许该请求。
		//AbortWithStatus：c.AbortWithStatus(200)的使用表示服务器在处理完OPTIONS请求后，将不会继续执行后续的中间件或路由处理程序，因为OPTIONS请求本身不包含实际数据，它只是一个检查。
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

// Timeout 超时中间件
func Timeout(t time.Duration) gin.HandlerFunc {
	// see:
	// https://github.com/vearne/gin-timeout
	// https://vearne.cc/archives/39135
	// https://github.com/gin-contrib/timeout
	return timeout.Timeout(
		timeout.WithTimeout(t),
	)
}
