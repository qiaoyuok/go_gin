package controllers

import (
	"github.com/gin-gonic/gin"
	"go_gin/dal/request"
	"go_gin/services"
	"go_gin/utils"
)

func UserLogin(c *gin.Context) {
	var req request.UserLoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResError(c, "请正确填写用户名-密码")
		return
	}

	if req.UserName != "admin" || req.Passwd != "123456" {
		utils.ResError(c, "用户名或密码不正确")
		return
	}

	token, err := services.GetToken(req.UserName, 1)
	if err != nil {
		utils.ResError(c, "生成Token失败")
		return
	}

	utils.ResSuccess(c, token)
	return
}
