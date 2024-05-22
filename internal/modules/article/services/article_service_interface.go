package services

import ArticleResponse "blog/internal/modules/article/reponses"

type ArticleRepositoryInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
}
