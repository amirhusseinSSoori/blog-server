package repositories

import ArtcleModel "blog/internal/modules/articles/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArtcleModel.Article
	Find(id int) ArtcleModel.Article
}
