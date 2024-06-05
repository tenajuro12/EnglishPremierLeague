package repository

import (
	"news_service/internal/entity"
)

type NewsArticleRepository interface {
	Save(article entity.NewsArticle) (entity.NewsArticle, error)
	FindByID(id int) (entity.NewsArticle, error)
	Update(article entity.NewsArticle) (entity.NewsArticle, error)
	Delete(id int) error
	FindAll() ([]entity.NewsArticle, error)
}
