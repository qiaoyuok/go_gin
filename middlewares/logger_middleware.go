package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_gin/utils"
	"io"
	"io/ioutil"
	"time"
)

// Logger 全局请求日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := parseBody(c)
		if err != nil {
			c.Abort()
			return
		}
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
		data["body"] = requestBody
		data["ip"] = c.ClientIP()
		data["user-agent"] = c.Request.UserAgent()
		data["cost"] = cost
		utils.ZapSugarLogger.Info(data)
	}
}

// parseBody 解析Body请求体
func parseBody(c *gin.Context) (bodyMap map[string]interface{}, err error) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	requestBody, _ := ioutil.ReadAll(tee)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	if len(requestBody) > 0 {
		err = json.Unmarshal(requestBody, &bodyMap)
	}
	return
}
