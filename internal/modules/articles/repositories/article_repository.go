package repositories

import (
	ArticleModel "blog/internal/modules/articles/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArtcileRepository struct {
	DB *gorm.DB
}

func New() *ArtcileRepository {
	return &ArtcileRepository{
		DB: database.Connection(),
	}

}

func (artcileRepository *ArtcileRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	artcileRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (artcileRepository *ArtcileRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article
	artcileRepository.DB.Joins("User").First(&article, id)
	return article
}
