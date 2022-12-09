package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/middlewares"
	"go_gin/routers"
	"go_gin/utils"
)

func main() {
	r := gin.New()
	r.Use(middlewares.Recovery(), middlewares.Logger())

	config.InitEnv()
	utils.RegisterLogger()

	routers.RegisterRouter(r)
	r.Run(":9300")
}
