package main

import (
	"EPL/match_service/internal/interfaces/handlers"
	"EPL/match_service/internal/interfaces/repository"
	"EPL/match_service/internal/usecases"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:AidosK05@localhost:5433/epl_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewMatchRepository(db)

	createUseCase := &usecases.CreateMatch{Repository: repo}
	getUseCase := &usecases.FindMatchByID{Repository: repo}
	updateUseCase := &usecases.UpdateMatch{Repository: repo}
	deleteUseCase := &usecases.DeleteMatch{Repository: repo}
	listUseCase := &usecases.FindAllMatches{Repository: repo}
	controller := &handlers.MatchHandler{
		CreateUseCase: createUseCase,
		GetUseCase:    getUseCase,
		UpdateUseCase: updateUseCase,
		DeleteUseCase: deleteUseCase,
		ListUseCase:   listUseCase}
	StartServer(controller)
}
