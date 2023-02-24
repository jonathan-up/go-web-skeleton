package logger

import (
	"github.com/gin-gonic/gin"
	"skeleton/bootstrap/logger"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()
		logger.SugarLogger.Debugf(
			"%d | %s | %s | %s | \"%s\"",
			ctx.Writer.Status(),
			end.Sub(start),
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.FullPath(),
		)
	}
}
