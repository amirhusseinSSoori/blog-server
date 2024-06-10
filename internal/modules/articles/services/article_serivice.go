package services

import (
	ArticleResponse "blog/internal/modules/articles/reponses"
	ArtcileRepository "blog/internal/modules/articles/repositories"
	"errors"
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

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.artcileRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return ArticleResponse.ToArticle(article), nil
}
