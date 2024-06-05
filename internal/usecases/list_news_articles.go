package usecases

import (
	"context"
	"news_service/internal/entity"
	"news_service/internal/interfaces/repository"
)

type ListNewsArticles struct {
	Repository repository.NewsArticleRepository
}

func (usecase *ListNewsArticles) Execute(ctx context.Context, u *ListNewsArticles) ([]*entity.NewsArticle, error) {
	articles, err := usecase.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var articlePointers []*entity.NewsArticle
	for i := range articles {
		articlePointers = append(articlePointers, &articles[i])
	}

	return articlePointers, nil
}
