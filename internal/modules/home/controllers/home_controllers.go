package controllers

import (
	ArtcileRepository "blog/internal/modules/article/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	artcileRepository ArtcileRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		artcileRepository: ArtcileRepository.New(),
	}

}

func (controller *Controller) Index(c *gin.Context) {
	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title": "HomePage",
	// })

	c.JSON(http.StatusOK, gin.H{
		"articles": controller.artcileRepository.List(8),
	})
}
