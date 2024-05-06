package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "modules/home/html/home", gin.H{
			"title": "HomePage",
		})

	})
}
