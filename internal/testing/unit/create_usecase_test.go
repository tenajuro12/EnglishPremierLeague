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

func TestCreateNewsArticle_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.CreateNewsArticle{
		Repository: mockRepo,
	}

	inputArticle := entity.NewsArticle{
		Title:    "Test Title",
		Content:  "Test Content",
		Category: "Test Category",
		Author:   "Test Author",
	}

	expectedArticle := inputArticle
	expectedArticle.ID = 1
	expectedArticle.CreatedAt = time.Now()
	expectedArticle.UpdatedAt = time.Now()

	mockRepo.On("Save", mock.AnythingOfType("entity.NewsArticle")).Return(expectedArticle, nil)

	resultArticle, err := useCase.Execute(inputArticle)

	assert.NoError(t, err)
	assert.Equal(t, expectedArticle.Title, resultArticle.Title)
	assert.Equal(t, expectedArticle.Content, resultArticle.Content)
	assert.Equal(t, expectedArticle.Category, resultArticle.Category)
	assert.Equal(t, expectedArticle.Author, resultArticle.Author)
	assert.WithinDuration(t, expectedArticle.CreatedAt, resultArticle.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedArticle.UpdatedAt, resultArticle.UpdatedAt, time.Second)

	mockRepo.AssertExpectations(t)
}

func TestCreateNewsArticle_Execute_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := &usecases.CreateNewsArticle{
		Repository: mockRepo,
	}

	inputArticle := entity.NewsArticle{
		Title:    "Test Title",
		Content:  "Test Content",
		Category: "Test Category",
		Author:   "Test Author",
	}

	mockRepo.On("Save", mock.AnythingOfType("entity.NewsArticle")).Return(entity.NewsArticle{}, errors.New("repository error"))

	resultArticle, err := useCase.Execute(inputArticle)

	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
	assert.Equal(t, entity.NewsArticle{}, resultArticle)

	mockRepo.AssertExpectations(t)
}
