package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(LoggerToFile())

	// 自定义错误处理
	r.Use(Recover())

	// header
	// NoCache is a middleware function that appends headers
	r.Use(NoCache())

	// 跨域处理
	r.Use(Cors())

	// Set X-Request-Id header
	r.Use(RequestId())
}
