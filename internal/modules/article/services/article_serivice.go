package services

import (
	ArticleResponse "blog/internal/modules/article/reponses"
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

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.artcileRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	article := articleService.artcileRepository.List(6)
	return ArticleResponse.ToArticles(article)
}
