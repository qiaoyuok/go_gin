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

func GetDYVideoList(c *gin.Context) {

	//services.GetDYVideoList("MS4wLjABAAAAUoce16bhp1iv971RZnOk0xcRNoxZ1gYAethlFJZyJhY", 0, 30)
	//services.GetDYVideoList("MS4wLjABAAAAKhBpJ9JK3NUJkF-mRZPOU_8AJdGRiR0EGMRYsH8O_vBbXURciLlYd1-_Vbn6JboT", 0, 30)
	services.GetDYVideoList("MS4wLjABAAAAyub-MX8KJ7rA875VbEuopiEvCRORS6aYAc7dfKy57cE", 0, 30)
}
