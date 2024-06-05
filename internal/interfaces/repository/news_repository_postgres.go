package repository

import (
	"database/sql"
	"news_service/internal/entity"
)

type NewsArticlePostgresRepository struct {
	db *sql.DB
}

func NewNewsArticlePostgresRepository(db *sql.DB) *NewsArticlePostgresRepository {
	return &NewsArticlePostgresRepository{db: db}
}

func (repo *NewsArticlePostgresRepository) Save(article entity.NewsArticle) (entity.NewsArticle, error) {
	query := `INSERT INTO news_articles (title, content, category, author, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := repo.db.QueryRow(query, article.Title, article.Content, article.Category, article.Author, article.CreatedAt, article.UpdatedAt).Scan(&article.ID)
	return article, err
}

func (repo *NewsArticlePostgresRepository) FindByID(id int) (entity.NewsArticle, error) {
	var article entity.NewsArticle
	query := `SELECT id, title, content, category, author, created_at, updated_at FROM news_articles WHERE id = $1`
	err := repo.db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Author, &article.CreatedAt, &article.UpdatedAt)
	return article, err
}

func (repo *NewsArticlePostgresRepository) Update(article entity.NewsArticle) (entity.NewsArticle, error) {
	query := `UPDATE news_articles SET title = $1, content = $2, category = $3, author = $4, updated_at = $5 WHERE id = $6`
	_, err := repo.db.Exec(query, article.Title, article.Content, article.Category, article.Author, article.UpdatedAt, article.ID)
	return article, err
}

func (repo *NewsArticlePostgresRepository) Delete(id int) error {
	query := `DELETE FROM news_articles WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *NewsArticlePostgresRepository) FindAll() ([]entity.NewsArticle, error) {
	query := `SELECT id, title, content, category, author, created_at, updated_at FROM news_articles`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []entity.NewsArticle
	for rows.Next() {
		var article entity.NewsArticle
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Author, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
