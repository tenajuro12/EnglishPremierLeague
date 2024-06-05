package handlers

import (
	"EPL/match_service/internal/entity"
	"EPL/match_service/internal/usecases"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MatchHandler struct {
	CreateUseCase *usecases.CreateMatch
	GetUseCase    *usecases.FindMatchByID
	UpdateUseCase *usecases.UpdateMatch
	DeleteUseCase *usecases.DeleteMatch
	ListUseCase   *usecases.FindAllMatches
}

func (h *MatchHandler) logRequest(r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)
}

func (h *MatchHandler) logError(err error) {
	log.Printf("Error: %v", err)
}

func (h *MatchHandler) logResponse(code int, response interface{}) {
	log.Printf("Responded with status code %d and body: %v", code, response)
}

func (h *MatchHandler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	var match entity.NewMatch
	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logError(err)
		return
	}

	createdMatch, err := h.CreateUseCase.Execute(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logError(err)
		return
	}

	response := map[string]interface{}{
		"message": "Match created successfully",
		"match":   createdMatch,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	h.logResponse(http.StatusCreated, response)
}

func (h *MatchHandler) GetMatch(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		h.logError(err)
		return
	}

	match, err := h.GetUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		h.logError(err)
		return
	}

	json.NewEncoder(w).Encode(match)

	h.logResponse(http.StatusOK, match)
}

func (h *MatchHandler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	var match entity.NewMatch
	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logError(err)
		return
	}

	updatedMatch, err := h.UpdateUseCase.Execute(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logError(err)
		return
	}

	response := map[string]interface{}{
		"message": "Match updated successfully",
		"match":   updatedMatch,
	}
	json.NewEncoder(w).Encode(response)

	h.logResponse(http.StatusOK, response)
}

func (h *MatchHandler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		h.logError(err)
		return
	}

	err = h.DeleteUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		h.logError(err)
		return
	}

	response := map[string]interface{}{
		"message": "Match deleted successfully",
	}
	json.NewEncoder(w).Encode(response)

	h.logResponse(http.StatusOK, response)
}

func (h *MatchHandler) ListMatches(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	matches, err := h.ListUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logError(err)
		return
	}

	json.NewEncoder(w).Encode(matches)

	h.logResponse(http.StatusOK, matches)
}

func (h *MatchHandler) UpcomingMatches(w http.ResponseWriter, r *http.Request) {
	h.logRequest(r)

	// Get all matches
	matches, err := h.ListUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logError(err)
		return
	}

	// Filter upcoming matches
	upcomingMatches := make([]entity.NewMatch, 0)
	now := time.Now()
	for _, match := range matches {
		matchDate, err := time.Parse(time.RFC3339, match.Date)
		if err != nil {
			continue
		}
		if matchDate.After(now) {
			upcomingMatches = append(upcomingMatches, match)
		}
	}

	sort.Slice(upcomingMatches, func(i, j int) bool {
		date1, _ := time.Parse(time.RFC3339, upcomingMatches[i].Date)
		date2, _ := time.Parse(time.RFC3339, upcomingMatches[j].Date)
		return date1.Before(date2)
	})

	if len(upcomingMatches) > 3 {
		upcomingMatches = upcomingMatches[:3]
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(upcomingMatches)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logError(err)
		return
	}

	h.logResponse(http.StatusOK, upcomingMatches)
}
