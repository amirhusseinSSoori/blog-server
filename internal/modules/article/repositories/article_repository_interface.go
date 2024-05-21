package repositories

import ArtcleModel "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArtcleModel.Article
}
