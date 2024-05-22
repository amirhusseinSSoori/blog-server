package services

import (
	ArticleModel "blog/internal/modules/article/models"
	ArtcileRepository "blog/internal/modules/article/repositories"
)

type ArticleService struct {
	artcileRepository ArtcileRepository.ArticleRepositoryInterface
}

func New() *ArticleService {

	return &ArticleService{
		ArtcileRepository.New(),
	}

}

func (articleService *ArticleService) GetFeaturedArticles() []ArticleModel.Article {
	return articleService.artcileRepository.List(4)
}

func (articleService *ArticleService) GetStoriesArticles() []ArticleModel.Article {
	return articleService.artcileRepository.List(6)
}
