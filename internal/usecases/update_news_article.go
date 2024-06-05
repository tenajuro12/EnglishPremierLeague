package usecases

import (
	"news_service/internal/entity"
	"news_service/internal/interfaces/repository"
	"time"
)

type UpdateNewsArticle struct {
	Repository repository.NewsArticleRepository
}

func (usecase *UpdateNewsArticle) Execute(article entity.NewsArticle) (entity.NewsArticle, error) {
	article.UpdatedAt = time.Now()
	return usecase.Repository.Update(article)
}
