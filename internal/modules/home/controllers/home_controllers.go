package controllers

import (
	ArticleService "blog/internal/modules/articles/services"
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}

}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "HomePage",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"featured": controller.articleService.GetFeaturedArticles(),
	// 	"stories":  controller.articleService.GetStoriesArticles(),
	// })
}
