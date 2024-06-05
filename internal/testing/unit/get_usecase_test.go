package repository_test

import (
	"errors"
	"news_service/internal/entity"
	"news_service/internal/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func TestGetNewsArticle_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.GetNewsArticle{
		Repository: mockRepo,
	}

	idToFetch := 1

	expectedArticle := entity.NewsArticle{
		ID:        idToFetch,
		Title:     "Test Title",
		Content:   "Test Content",
		Category:  "Test Category",
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByID", idToFetch).Return(expectedArticle, nil)

	result, err := useCase.Execute(idToFetch)

	assert.NoError(t, err)
	assert.Equal(t, expectedArticle, result)

	mockRepo.AssertCalled(t, "FindByID", idToFetch)
	mockRepo.AssertExpectations(t)
}

func TestGetNewsArticle_Execute_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.GetNewsArticle{
		Repository: mockRepo,
	}

	idToFetch := 1

	mockRepo.On("FindByID", idToFetch).Return(entity.NewsArticle{}, errors.New("repository error"))

	result, err := useCase.Execute(idToFetch)

	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
	assert.Equal(t, entity.NewsArticle{}, result)

	mockRepo.AssertCalled(t, "FindByID", idToFetch)
	mockRepo.AssertExpectations(t)
}

func (m *MockRepository) Save(article entity.NewsArticle) (entity.NewsArticle, error) {
	args := m.Called(article)
	return args.Get(0).(entity.NewsArticle), args.Error(1)
}

func (m *MockRepository) FindByID(id int) (entity.NewsArticle, error) {
	args := m.Called(id)
	return args.Get(0).(entity.NewsArticle), args.Error(1)
}

func (m *MockRepository) Update(article entity.NewsArticle) (entity.NewsArticle, error) {
	args := m.Called(article)
	return args.Get(0).(entity.NewsArticle), args.Error(1)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) FindAll() ([]entity.NewsArticle, error) {
	args := m.Called()
	return args.Get(0).([]entity.NewsArticle), args.Error(1)
}
