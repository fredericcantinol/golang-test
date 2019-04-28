package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"./Controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/player", Controller.NewPlayer).Methods("POST")
	r.HandleFunc("/player", Controller.DeletePlayer).Methods("DELETE")
	r.HandleFunc("/player", Controller.SeePlayer).Methods("GET")
	r.HandleFunc("/players", Controller.SeePlayers).Methods("GET")
	r.HandleFunc("/game", Controller.NewGame).Methods("POST")
	r.HandleFunc("/game", Controller.DeleteGame).Methods("DELETE")
	r.HandleFunc("/game", Controller.SeeGame).Methods("GET")
	r.HandleFunc("/games", Controller.SeeGames).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}