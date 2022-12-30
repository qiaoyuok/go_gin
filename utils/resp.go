package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResSuccess 成功返回
func ResSuccess(c *gin.Context, data interface{}) {
	if _, ok := data.(string); ok {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    data,
		})
	}
	return
}

// ResError 错误返回
func ResError(c *gin.Context, msg interface{}) {
	res := gin.H{
		"status": 500,
		"msg":    msg,
	}
	ZapSugarLogger.Error(res)
	c.JSON(http.StatusOK, res)
	return
}
