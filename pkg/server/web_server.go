package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"news_service/internal/interfaces/handlers"
)

func StartServer(handler *handlers.NewsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/articles", handler.CreateNewsArticleHandler).Methods("POST")
	router.HandleFunc("/articles/{id:[0-9]+}", handler.GetNewsArticle).Methods("GET")
	router.HandleFunc("/articles/{id:[0-9]+}", handler.UpdateNewsArticle).Methods("PUT")
	router.HandleFunc("/articles/{id:[0-9]+}", handler.DeleteNewsArticle).Methods("DELETE")
	router.HandleFunc("/articles", handler.ListNewsArticles).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
