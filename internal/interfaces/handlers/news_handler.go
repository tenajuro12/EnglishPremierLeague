package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"news_service/internal/entity"
	"news_service/internal/usecases"
	"strconv"
)

type NewsHandler struct {
	CreateUseCase *usecases.CreateNewsArticle
	GetUseCase    *usecases.GetNewsArticle
	UpdateUseCase *usecases.UpdateNewsArticle
	DeleteUseCase *usecases.DeleteNewsArticle
	ListUseCase   *usecases.ListNewsArticles
}

func (c *NewsHandler) CreateNewsArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article entity.NewsArticle
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle, err := c.CreateUseCase.Execute(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdArticle)
}

func (controller *NewsHandler) GetNewsArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid article ID: %v", err)
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := controller.GetUseCase.Execute(id)
	if err != nil {
		log.Printf("Error retrieving news article: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func (controller *NewsHandler) UpdateNewsArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid article ID: %v", err)
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var article entity.NewsArticle
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article.ID = id
	updatedArticle, err := controller.UpdateUseCase.Execute(article)
	if err != nil {
		log.Printf("Error updating news article: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedArticle)
}

func (controller *NewsHandler) DeleteNewsArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid article ID: %v", err)
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	if err := controller.DeleteUseCase.Execute(id); err != nil {
		log.Printf("Error deleting news article: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *NewsHandler) ListNewsArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := controller.ListUseCase.Execute(nil, nil)
	if err != nil {
		log.Printf("Error listing news articles: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
