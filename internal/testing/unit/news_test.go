package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"news_service/internal/entity"
	"news_service/internal/interfaces/repository"
	"testing"
	"time"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open sqlmock database connection: %v", err)
	}

	cleanup := func() {
		db.Close()
	}
	return db, mock, cleanup
}

func TestSave(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := repository.NewNewsArticlePostgresRepository(db)

	article := entity.NewsArticle{
		Title:     "Test Title",
		Content:   "Test Content",
		Category:  "Test Category",
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `INSERT INTO news_articles \(title, content, category, author, created_at, updated_at\)
              VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\) RETURNING id`

	mock.ExpectQuery(query).
		WithArgs(article.Title, article.Content, article.Category, article.Author, article.CreatedAt, article.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	resultArticle, err := repo.Save(article)
	assert.NoError(t, err)
	assert.Equal(t, 1, resultArticle.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := repository.NewNewsArticlePostgresRepository(db)

	query := `SELECT id, title, content, category, author, created_at, updated_at FROM news_articles WHERE id = \$1`

	article := entity.NewsArticle{
		ID:        1,
		Title:     "Test Title",
		Content:   "Test Content",
		Category:  "Test Category",
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "category", "author", "created_at", "updated_at"}).
			AddRow(article.ID, article.Title, article.Content, article.Category, article.Author, article.CreatedAt, article.UpdatedAt))

	result, err := repo.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, article, result)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := repository.NewNewsArticlePostgresRepository(db)

	article := entity.NewsArticle{
		ID:        1,
		Title:     "Test Title",
		Content:   "Test Content",
		Category:  "Test Category",
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `UPDATE news_articles SET title = \$1, content = \$2, category = \$3, author = \$4, updated_at = \$5 WHERE id = \$6`

	mock.ExpectExec(query).
		WithArgs(article.Title, article.Content, article.Category, article.Author, article.UpdatedAt, article.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.Update(article)
	assert.NoError(t, err)
	assert.Equal(t, article, result)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := repository.NewNewsArticlePostgresRepository(db)

	query := `DELETE FROM news_articles WHERE id = \$1`

	mock.ExpectExec(query).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := repository.NewNewsArticlePostgresRepository(db)

	query := `SELECT id, title, content, category, author, created_at, updated_at FROM news_articles`

	match1 := entity.NewsArticle{
		ID:        1,
		Title:     "Title 1",
		Content:   "Content 1",
		Category:  "Category 1",
		Author:    "Author 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	match2 := entity.NewsArticle{
		ID:        2,
		Title:     "Title 2",
		Content:   "Content 2",
		Category:  "Category 2",
		Author:    "Author 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(query).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "category", "author", "created_at", "updated_at"}).
			AddRow(match1.ID, match1.Title, match1.Content, match1.Category, match1.Author, match1.CreatedAt, match1.UpdatedAt).
			AddRow(match2.ID, match2.Title, match2.Content, match2.Category, match2.Author, match2.CreatedAt, match2.UpdatedAt))

	results, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, match1, results[0])
	assert.Equal(t, match2, results[1])

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
