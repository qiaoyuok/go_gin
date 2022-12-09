package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/dal/biz"
	"go_gin/dal/request"
	"net/http"
)

func Index(c *gin.Context) {

	idStr := c.Param("id")
	fmt.Printf("%#v\nid:%s\n", config.C.Log, idStr)
	users, _ := biz.ListUser()
	c.JSON(http.StatusOK, users)
}

func Create(c *gin.Context) {
	var req request.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(111)
	}
	fmt.Printf("%#v\n", req)

	biz.Create(req)

	c.JSON(http.StatusOK, req)
}
