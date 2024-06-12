package controllers

import (
	ArticleService "blog/internal/modules/articles/services"
	"blog/pkg/html"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleSerivceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}

}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "server Error", "message": "error converting the id"})
		return
	}

	// Find the article from the dataBase
	articel, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/404", gin.H{"title": "page not found", "message": err.Error()})
		return
	}

	//If article found, render artilce temaplate
	html.Render(c, http.StatusOK, "modules/articles/html/show", gin.H{"title": "show Article", "article": articel})
}
