// usecases/get_news_article.go

package usecases

import (
	"news_service/internal/entity"
	"news_service/internal/interfaces/repository"
)

type GetNewsArticle struct {
	Repository repository.NewsArticleRepository
}

func (usecase *GetNewsArticle) Execute(id int) (entity.NewsArticle, error) {
	return usecase.Repository.FindByID(id)
}
