package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/dal/biz"
	"go_gin/dal/request"
	"go_gin/services"
	"go_gin/utils"
	"net/http"
	"sync"
)

func UserList(c *gin.Context) {

	idStr := c.Param("id")
	fmt.Printf("%#v\nid:%s\n", config.C.Log, idStr)
	users, _ := biz.ListUser()
	panic("娃哈哈")
	c.JSON(http.StatusOK, users)
}

func UserCreate(c *gin.Context) {
	var req request.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(111)
	}
	fmt.Printf("%#v\n", req)

	biz.Create(req)

	c.JSON(http.StatusOK, req)
}

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

func BdSearch(c *gin.Context) {
	wg := sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(page int) {
			services.SearchBd(&wg, "最新电影", page)
		}(i)
	}

	wg.Wait()
	c.String(http.StatusOK, "ok")
}
