package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_gin/utils"
	"net/http"
)

// Recovery 全局异常捕获中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ZapSugarLogger.Error(r)
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  err2String(r),
					"data": nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// err2String 错误转字符串
func err2String(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
