package repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"news_service/internal/entity"
	"news_service/internal/usecases"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func TestDeleteNewsArticle_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.DeleteNewsArticle{
		Repository: mockRepo,
	}

	idToDelete := 1

	mockRepo.On("Delete", idToDelete).Return(nil)

	err := useCase.Execute(idToDelete)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "Delete", idToDelete)
	mockRepo.AssertExpectations(t)
}

func TestDeleteNewsArticle_Execute_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.DeleteNewsArticle{
		Repository: mockRepo,
	}

	idToDelete := 1

	mockRepo.On("Delete", idToDelete).Return(errors.New("repository error"))

	err := useCase.Execute(idToDelete)

	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")

	mockRepo.AssertCalled(t, "Delete", idToDelete)
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
