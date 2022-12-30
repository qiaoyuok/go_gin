package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_gin/services"
	"go_gin/utils"
	"strings"
)

// CheckAuth 校验登录
func CheckAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authBearerHeader := c.Request.Header.Get("Authorization")
		token := c.Request.Header.Get("Token")
		if len(authBearerHeader) == 0 && len(token) == 0 {
			utils.ResError(c, "系统出错")
			c.Abort()
			return
		}
		// 按空格分割
		if len(authBearerHeader) != 0 {
			parts := strings.SplitN(authBearerHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				utils.ResError(c, "系统出错")
				c.Abort()
				return
			}
			token = parts[1]
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := services.ParseToken(token)
		if err != nil {
			utils.ResError(c, "系统出错")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userName", mc.UserName)
		c.Set("ID", mc.ID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
