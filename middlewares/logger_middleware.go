package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_gin/utils"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		data := make(map[string]interface{}, 0)
		data["path"] = path
		data["status"] = c.Writer.Status()
		data["method"] = c.Request.Method
		data["path"] = path
		data["query"] = query
		data["ip"] = c.ClientIP()
		data["user-agent"] = c.Request.UserAgent()
		data["cost"] = cost
		utils.ZapSugarLogger.Info(data)
	}
}
