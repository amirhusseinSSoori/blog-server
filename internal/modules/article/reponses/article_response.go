package reponses

import (
	ArticleModel "blog/internal/modules/article/models"
	UserResponse "blog/internal/modules/user/reponses"
	"fmt"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
