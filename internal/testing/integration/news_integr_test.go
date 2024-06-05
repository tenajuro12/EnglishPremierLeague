package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"news_service/internal/entity"
	"news_service/internal/interfaces/handlers"
	"testing"
)

type MockNewsUseCase struct{}

func TestNewsIntegration(t *testing.T) {
	mockUseCase := &MockNewsUseCase{}
	handler := &handlers.NewsHandler{
		CreateUseCase: mockUseCase,
		GetUseCase:    mockUseCase,
		UpdateUseCase: mockUseCase,
		DeleteUseCase: mockUseCase,
		ListUseCase:   mockUseCase,
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/articles":
			switch r.Method {
			case http.MethodPost:
				handler.CreateNewsArticleHandler(w, r)
			case http.MethodGet:
				handler.ListNewsArticles(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	defer testServer.Close()

	// Test creating a news article
	article := entity.NewsArticle{Title: "Test Article", Content: "Test Content"}
	articleJSON, _ := json.Marshal(article)
	resp, _ := http.Post(testServer.URL+"/articles", "application/json", bytes.NewBuffer(articleJSON))
	resp.Body.Close()

	// Test listing news articles
	resp, _ = http.Get(testServer.URL + "/articles")
	defer resp.Body.Close()

	var articles []entity.NewsArticle
	json.NewDecoder(resp.Body).Decode(&articles)

	expectedArticles := []entity.NewsArticle{
		{ID: 1, Title: "Article 1", Content: "Content 1"},
		{ID: 2, Title: "Article 2", Content: "Content 2"},
	}

	for i, expected := range expectedArticles {
		if articles[i] != expected {
			t.Errorf("Expected article %v but got %v", expected, articles[i])
		}
	}
}

func (uc *MockGetNewsUseCase) Execute(id int) (entity.NewsArticle, error) {
	return entity.NewsArticle{ID: id, Title: "Mock Title", Content: "Mock Content"}, nil
}

func TestGetNewsArticle(t *testing.T) {
	mockUseCase := &MockGetNewsUseCase{}
	handler := &handlers.NewsHandler{
		GetUseCase: mockUseCase,
	}

	req := httptest.NewRequest("GET", "/articles/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", handler.GetNewsArticle)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var article entity.NewsArticle
	err := json.Unmarshal(w.Body.Bytes(), &article)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedArticle := entity.NewsArticle{ID: 1, Title: "Mock Title", Content: "Mock Content"}
	if article != expectedArticle {
		t.Errorf("Expected article %+v but got %+v", expectedArticle, article)
	}
}

func (uc *MockGetNewsUseCase) Execute(id int) (entity.NewsArticle, error) {
	return entity.NewsArticle{ID: id, Title: "Mock Title", Content: "Mock Content"}, nil
}

func TestGetNewsArticle(t *testing.T) {
	mockUseCase := &MockGetNewsUseCase{}
	handler := &handlers.NewsHandler{
		GetUseCase: mockUseCase,
	}

	req := httptest.NewRequest("GET", "/articles/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", handler.GetNewsArticle)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var article entity.NewsArticle
	err := json.Unmarshal(w.Body.Bytes(), &article)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedArticle := entity.NewsArticle{ID: 1, Title: "Mock Title", Content: "Mock Content"}
	if article != expectedArticle {
		t.Errorf("Expected article %+v but got %+v", expectedArticle, article)
	}
}

type MockUpdateNewsUseCase struct{}

func (uc *MockUpdateNewsUseCase) Execute(article entity.NewsArticle) (entity.NewsArticle, error) {
	return article, nil
}

type MockDeleteNewsUseCase struct{}

func (uc *MockDeleteNewsUseCase) Execute(id int) error {
	return nil
}

type MockListNewsUseCase struct{}

func (uc *MockListNewsUseCase) Execute(filters map[string]interface{}, pagination *entity.Pagination) ([]entity.NewsArticle, error) {
	return []entity.NewsArticle{
		{ID: 1, Title: "Article 1", Content: "Content 1"},
		{ID: 2, Title: "Article 2", Content: "Content 2"},
	}, nil
}

func TestUpdateNewsArticle(t *testing.T) {
	mockUseCase := &MockUpdateNewsUseCase{}
	handler := &handlers.NewsHandler{
		UpdateUseCase: mockUseCase,
	}

	reqBody := `{"id":1,"title":"Updated Article","content":"Updated Content"}`
	req := httptest.NewRequest("PUT", "/articles/1", nil)
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(reqBody)))
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", handler.UpdateNewsArticle)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var updatedArticle entity.NewsArticle
	err := json.Unmarshal(w.Body.Bytes(), &updatedArticle)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedArticle := entity.NewsArticle{ID: 1, Title: "Updated Article", Content: "Updated Content"}
	if updatedArticle != expectedArticle {
		t.Errorf("Expected article %+v but got %+v", expectedArticle, updatedArticle)
	}
}

func TestDeleteNewsArticle(t *testing.T) {
	mockUseCase := &MockDeleteNewsUseCase{}
	handler := &handlers.NewsHandler{
		DeleteUseCase: mockUseCase,
	}

	req := httptest.NewRequest("DELETE", "/articles/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", handler.DeleteNewsArticle)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got %d", http.StatusNoContent, w.Code)
	}
}

func TestListNewsArticles(t *testing.T) {
	mockUseCase := &MockListNewsUseCase{}
	handler := &handlers.NewsHandler{
		ListUseCase: mockUseCase,
	}

	req := httptest.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.ListNewsArticles)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var articles []entity.NewsArticle
	err := json.Unmarshal(w.Body.Bytes(), &articles)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedArticles := []entity.NewsArticle{
		{ID: 1, Title: "Article 1", Content: "Content 1"},
		{ID: 2, Title: "Article 2", Content: "Content 2"},
	}
	if len(articles) != len(expectedArticles) {
		t.Errorf("Expected %d articles but got %d", len(expectedArticles), len(articles))
	}

	for i, article := range articles {
		if article != expectedArticles[i] {
			t.Errorf("Expected article %+v but got %+v", expectedArticles[i], article)
		}
	}
}
func (uc *MockNewsUseCase) Execute(article entity.NewsArticle) (entity.NewsArticle, error) {
	return entity.NewsArticle{ID: 1, Title: article.Title, Content: article.Content}, nil
}

func (uc *MockNewsUseCase) ExecuteWithID(id int) (entity.NewsArticle, error) {
	return entity.NewsArticle{ID: id, Title: "Mock Title", Content: "Mock Content"}, nil
}

func (uc *MockNewsUseCase) Update(article entity.NewsArticle) (entity.NewsArticle, error) {
	return article, nil
}

func (uc *MockNewsUseCase) Delete(id int) error {
	return nil
}

func (uc *MockNewsUseCase) List() ([]entity.NewsArticle, error) {
	return []entity.NewsArticle{
		{ID: 1, Title: "Article 1", Content: "Content 1"},
		{ID: 2, Title: "Article 2", Content: "Content 2"},
	}, nil
}
