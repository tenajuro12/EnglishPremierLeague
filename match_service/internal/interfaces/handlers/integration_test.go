package handlers_test

import (
	"EPL/match_service/internal/entity"
	"EPL/match_service/internal/interfaces/handlers"
	"EPL/match_service/internal/usecases"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setup() (handlers.MatchHandler, *mux.Router) {
	router := mux.NewRouter()
	createUseCase := &usecases.CreateMatch{}
	getUseCase := &usecases.FindMatchByID{}
	updateUseCase := &usecases.UpdateMatch{}
	deleteUseCase := &usecases.DeleteMatch{}
	listUseCase := &usecases.FindAllMatches{}
	handler := handlers.MatchHandler{
		CreateUseCase: createUseCase,
		GetUseCase:    getUseCase,
		UpdateUseCase: updateUseCase,
		DeleteUseCase: deleteUseCase,
		ListUseCase:   listUseCase,
	}
	return handler, router
}

func TestCreateMatch(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/match", handler.CreateMatch).Methods("POST")

	match := entity.NewMatch{
		HomeTeam: "Team A",
		AwayTeam: "Team B",
		Date:     time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}

	body, _ := json.Marshal(match)
	req, err := http.NewRequest("POST", "/match", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Match created successfully", response["message"])
}

func TestGetMatch(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/match/{id}", handler.GetMatch).Methods("GET")

	req, err := http.NewRequest("GET", "/match/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var match entity.NewMatch
	err = json.NewDecoder(rr.Body).Decode(&match)
	assert.NoError(t, err)
	assert.Equal(t, 1, match.ID)
	assert.Equal(t, "Team A", match.HomeTeam)
	assert.Equal(t, "Team B", match.AwayTeam)
}

func TestUpdateMatch(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/match/{id}", handler.UpdateMatch).Methods("PUT")

	match := entity.NewMatch{
		ID:       1,
		HomeTeam: "Team A Updated",
		AwayTeam: "Team B Updated",
		Date:     time.Now().Add(48 * time.Hour).Format(time.RFC3339),
	}

	body, _ := json.Marshal(match)
	req, err := http.NewRequest("PUT", "/match/1", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Match updated successfully", response["message"])
}

func TestDeleteMatch(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/match/{id}", handler.DeleteMatch).Methods("DELETE")

	req, err := http.NewRequest("DELETE", "/match/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Match deleted successfully", response["message"])
}

func TestListMatches(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/matches", handler.ListMatches).Methods("GET")

	req, err := http.NewRequest("GET", "/matches", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var matches []entity.NewMatch
	err = json.NewDecoder(rr.Body).Decode(&matches)
	assert.NoError(t, err)
	assert.Greater(t, len(matches), 0)
}

func TestUpcomingMatches(t *testing.T) {
	handler, router := setup()
	router.HandleFunc("/upcoming_matches", handler.UpcomingMatches).Methods("GET")

	req, err := http.NewRequest("GET", "/upcoming_matches", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var upcomingMatches []entity.NewMatch
	err = json.NewDecoder(rr.Body).Decode(&upcomingMatches)
	assert.NoError(t, err)
	assert.LessOrEqual(t, len(upcomingMatches), 3)
	for _, match := range upcomingMatches {
		matchDate, _ := time.Parse(time.RFC3339, match.Date)
		assert.True(t, matchDate.After(time.Now()))
	}
}
