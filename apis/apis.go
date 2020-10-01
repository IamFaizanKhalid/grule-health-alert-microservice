package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../constants"
	"../data"
)

// Home is handler for /.
func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v]\t%s", r.Method, constants.HomeAPIRoute)
	fmt.Fprintf(w, "API seems fine.")
}

// GetAllPeople is handler for /text.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v]\t%s", r.Method, constants.GetAllPeopleAPIRoute)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&data.People); err != nil {
		log.Printf("[ERROR]\t%v", err)
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}
}
