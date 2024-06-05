package main

import (
	"EPL/match_service/internal/interfaces/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer(handler *handlers.MatchHandler) {
	router := mux.NewRouter()
	router.HandleFunc("/matches", handler.CreateMatch).Methods("POST")
	router.HandleFunc("/matches/{id:[0-9]+}", handler.GetMatch).Methods("GET")
	router.HandleFunc("/matches/{id:[0-9]+}", handler.UpdateMatch).Methods("PUT")
	router.HandleFunc("/matches/{id:[0-9]+}", handler.DeleteMatch).Methods("DELETE")
	router.HandleFunc("/matches", handler.ListMatches).Methods("GET")
	router.HandleFunc("/upcoming/matches", handler.UpcomingMatches).Methods("GET")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
