package usecases

import "news_service/internal/interfaces/repository"

type DeleteNewsArticle struct {
	Repository repository.NewsArticleRepository
}

func (usecase *DeleteNewsArticle) Execute(id int) error {
	return usecase.Repository.Delete(id)
}
