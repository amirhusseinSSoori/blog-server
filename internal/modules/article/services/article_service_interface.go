package services

import articleModel "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	GetFeaturedArticles() []articleModel.Article
	GetStoriesArticles() []articleModel.Article
}
