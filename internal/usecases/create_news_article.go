package usecases

import (
	"news_service/internal/entity"
	"news_service/internal/interfaces/repository"
	"time"
)

type CreateNewsArticle struct {
	Repository repository.NewsArticleRepository
}

func (uc *CreateNewsArticle) Execute(article entity.NewsArticle) (entity.NewsArticle, error) {
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	return uc.Repository.Save(article)
}
