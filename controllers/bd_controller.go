package controllers

import (
	"github.com/gin-gonic/gin"
	"go_gin/services"
	"net/http"
	"sync"
)

func BdSearch(c *gin.Context) {
	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(page int) {
			services.SearchBd(&wg, "酒店设备", page)
		}(i)
	}

	wg.Wait()
	c.String(http.StatusOK, "ok")
}

func GetArticleTKD(c *gin.Context) {
	services.GetTKD("https://www.poluoa.com/28855/")
}
