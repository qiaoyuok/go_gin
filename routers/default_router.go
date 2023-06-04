package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin/controllers"
)

func RegisterRouter(r *gin.Engine) {

	//au := r.Group("", middlewares.CheckAuth())

	r.POST("/user_login", controllers.UserLogin)
	r.GET("/bd-search", controllers.BdSearch)
	r.GET("/article-search", controllers.GetArticleTKD)
	r.GET("/dy-download", controllers.GetDYVideoList)

}
