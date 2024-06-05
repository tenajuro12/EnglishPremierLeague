package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"news_service/internal/interfaces/handlers"
	"news_service/internal/interfaces/repository"
	"news_service/internal/usecases"
	"news_service/pkg/server"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/news?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewNewsArticlePostgresRepository(db)

	createUseCase := &usecases.CreateNewsArticle{Repository: repo}
	getUseCase := &usecases.GetNewsArticle{Repository: repo}
	updateUseCase := &usecases.UpdateNewsArticle{Repository: repo}
	deleteUseCase := &usecases.DeleteNewsArticle{Repository: repo}
	listUseCase := &usecases.ListNewsArticles{Repository: repo}

	controller := &handlers.NewsHandler{
		CreateUseCase: createUseCase,
		GetUseCase:    getUseCase,
		UpdateUseCase: updateUseCase,
		DeleteUseCase: deleteUseCase,
		ListUseCase:   listUseCase,
	}

	server.StartServer(controller)
}
