package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./apis"
	"./constants"
)

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc(constants.HomeAPIRoute, apis.Home).Methods("GET")
	r.HandleFunc(constants.GetAllPeopleAPIRoute, apis.GetAllPeople).Methods("GET")
	r.HandleFunc(constants.UpdatePersonAPIRoute, apis.UpdatePerson).Methods("PATCH")

	// Config
	r.Host("*")
	r.Use(mux.CORSMethodMiddleware(r))

	//Run
	log.Printf("[START]\tRunning server at %s ...\n", constants.GetAPIAddress())
	log.Printf("[ERROR]\t%v", http.ListenAndServe(constants.GetHTTPPort(), r))
}
