package services

import ArticleResponse "blog/internal/modules/articles/reponses"

type ArticleSerivceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
}
