package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin/controllers"
)

func RegisterRouter(r *gin.Engine) {

	//au := r.Group("", middlewares.CheckAuth())

	r.POST("/user_login", controllers.UserLogin)
	r.GET("/bd-search", controllers.BdSearch)

}
