package routes

import (
	articleCtrl "blog/internal/modules/articles/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articleCtrl.New()
	router.GET("/articles/:id", articlesController.Show)
}
