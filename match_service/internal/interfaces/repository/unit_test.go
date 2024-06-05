package repository

import (
	"EPL/match_service/internal/entity"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

func TestInsert(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewMatchRepository(db)

	match := entity.NewMatch{
		HomeTeam:  "Team A",
		AwayTeam:  "Team B",
		Date:      "2024-06-01",
		Status:    "Scheduled",
		HomeScore: 0,
		AwayScore: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `INSERT INTO matches \(home_team, away_team, date, status, home_score, away_score, created_at, updated_at\)
              VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\) RETURNING id`

	mock.ExpectQuery(query).
		WithArgs(match.HomeTeam, match.AwayTeam, match.Date, match.Status, match.HomeScore, match.AwayScore, match.CreatedAt, match.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := repo.Insert(match)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewMatchRepository(db)

	query := `SELECT id, home_team, away_team, date, status, home_score, away_score, created_at, updated_at FROM matches WHERE id = \$1`

	match := entity.NewMatch{
		ID:        1,
		HomeTeam:  "Team A",
		AwayTeam:  "Team B",
		Date:      "2024-06-01",
		Status:    "Scheduled",
		HomeScore: 0,
		AwayScore: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "home_team", "away_team", "date", "status", "home_score", "away_score", "created_at", "updated_at"}).
			AddRow(match.ID, match.HomeTeam, match.AwayTeam, match.Date, match.Status, match.HomeScore, match.AwayScore, match.CreatedAt, match.UpdatedAt))

	result, err := repo.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, match, result)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewMatchRepository(db)

	match := entity.NewMatch{
		ID:        1,
		HomeTeam:  "Team A",
		AwayTeam:  "Team B",
		Date:      "2024-06-01",
		Status:    "Completed",
		HomeScore: 2,
		AwayScore: 1,
		UpdatedAt: time.Now(),
	}

	query := `UPDATE matches SET home_team = \$1, away_team = \$2, date = \$3, status = \$4, home_score = \$5, away_score = \$6, updated_at = \$7 WHERE id = \$8`

	mock.ExpectExec(query).
		WithArgs(match.HomeTeam, match.AwayTeam, match.Date, match.Status, match.HomeScore, match.AwayScore, match.UpdatedAt, match.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.Update(match)
	assert.NoError(t, err)
	assert.Equal(t, match, result)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewMatchRepository(db)

	query := `DELETE FROM matches WHERE id = \$1`

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

	repo := NewMatchRepository(db)

	query := `SELECT id, home_team, away_team, date, status, home_score, away_score, created_at, updated_at FROM matches`

	match1 := entity.NewMatch{
		ID:        1,
		HomeTeam:  "Team A",
		AwayTeam:  "Team B",
		Date:      "2024-06-01",
		Status:    "Scheduled",
		HomeScore: 0,
		AwayScore: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	match2 := entity.NewMatch{
		ID:        2,
		HomeTeam:  "Team C",
		AwayTeam:  "Team D",
		Date:      "2024-06-02",
		Status:    "Scheduled",
		HomeScore: 0,
		AwayScore: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(query).
		WillReturnRows(sqlmock.NewRows([]string{"id", "home_team", "away_team", "date", "status", "home_score", "away_score", "created_at", "updated_at"}).
			AddRow(match1.ID, match1.HomeTeam, match1.AwayTeam, match1.Date, match1.Status, match1.HomeScore, match1.AwayScore, match1.CreatedAt, match1.UpdatedAt).
			AddRow(match2.ID, match2.HomeTeam, match2.AwayTeam, match2.Date, match2.Status, match2.HomeScore, match2.AwayScore, match2.CreatedAt, match2.UpdatedAt))

	results, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, match1, results[0])
	assert.Equal(t, match2, results[1])

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
