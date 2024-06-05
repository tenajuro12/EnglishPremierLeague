package repository

import (
	"EPL/match_service/internal/entity"
	"database/sql"
	"errors"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func (repo *MatchRepository) Insert(match entity.NewMatch) (entity.NewMatch, error) {
	query := `INSERT INTO matches (home_team, away_team, date, status, home_score, away_score, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := repo.db.QueryRow(query, match.HomeTeam, match.AwayTeam, match.Date, match.Status, match.HomeScore, match.AwayScore, match.CreatedAt, match.UpdatedAt).Scan(&match.ID)
	return match, err
}

func (repo *MatchRepository) FindByID(id int) (entity.NewMatch, error) {
	var match entity.NewMatch
	query := `SELECT id, home_team, away_team, date, status, home_score, away_score, created_at, updated_at FROM matches WHERE id = $1`
	err := repo.db.QueryRow(query, id).Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.Date, &match.Status, &match.HomeScore, &match.AwayScore, &match.CreatedAt, &match.UpdatedAt)
	if err == sql.ErrNoRows {
		return match, ErrRecordNotFound
	}
	return match, err
}

func (repo *MatchRepository) Update(match entity.NewMatch) (entity.NewMatch, error) {
	query := `UPDATE matches SET home_team = $1, away_team = $2, date = $3, status = $4, home_score = $5, away_score = $6, updated_at = $7 WHERE id = $8`
	_, err := repo.db.Exec(query, match.HomeTeam, match.AwayTeam, match.Date, match.Status, match.HomeScore, match.AwayScore, match.UpdatedAt, match.ID)
	return match, err
}

func (repo *MatchRepository) Delete(id int) error {
	query := `DELETE FROM matches WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *MatchRepository) GetAll() ([]entity.NewMatch, error) {
	query := `SELECT id, home_team, away_team, date, status, home_score, away_score, created_at, updated_at FROM matches`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []entity.NewMatch
	for rows.Next() {
		var match entity.NewMatch
		err := rows.Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.Date, &match.Status, &match.HomeScore, &match.AwayScore, &match.CreatedAt, &match.UpdatedAt)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}
