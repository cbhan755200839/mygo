package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/cbhan755200839/mygo/utils/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogMiddleWare(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求方法
		reqMethod := c.Request.Method
		// 请求uri
		reqURI := c.Request.RequestURI
		// 状态
		statusCode := c.Writer.Status()
		if reqURI == "/favicon.ico" ||
			reqURI == "/.well-known/appspecific/com.chrome.devtools.json" ||
			strings.HasPrefix(reqURI, "/static") {
			return
		}
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 请求处理时间
		duration := time.Since(startTime)

		if statusCode == http.StatusOK {
			logs.Log.Info("HTTP请求成功",
				zap.String("method", reqMethod),
				zap.String("uri", reqURI),
				zap.Int("statusCode", statusCode),
				zap.Duration("duration", duration),
			)
		} else {
			logs.Log.Error("HTTP请求异常",
				zap.String("method", reqMethod),
				zap.String("uri", reqURI),
				zap.Int("statusCode", statusCode),
				zap.Duration("duration", duration),
			)
		}
	}
}
