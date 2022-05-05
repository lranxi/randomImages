package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func LogMiddleware(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()

		defer func() {
			latency := time.Since(start)
			clientIP := c.ClientIP()
			clientUserAgent := c.Request.UserAgent()

			log.Info("access",
				zap.Any("method", c.Request.Method),
				zap.Any("path", path),
				zap.Any("clientIP", clientIP),
				zap.Any("clientUserAgent", clientUserAgent),
				zap.Any("http_code", c.Writer.Status()),
				zap.Any("success", c.Writer.Status() == http.StatusOK),
				zap.Any("latency", latency),
			)
		}()

		c.Next()
	}
}
