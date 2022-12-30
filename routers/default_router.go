package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin/controllers"
	"go_gin/middlewares"
)

func RegisterRouter(r *gin.Engine) {

	au := r.Group("", middlewares.CheckAuth())

	au.POST("/user-create", controllers.UserCreate)
	au.GET("/user-list", controllers.UserList)
	r.POST("/user_login", controllers.UserLogin)

}
