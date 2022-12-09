package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin/controllers"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/:id", controllers.Index)
	r.POST("/create", controllers.Create)

}
