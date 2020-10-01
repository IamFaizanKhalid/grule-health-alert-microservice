package apis

import (
	"../constants"
	"../data"
	"../grules"
	"../models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var dataCtx = grules.GetDataContext()
var knowledgeBase = grules.GetKnowledgeBase()

// Home is handler for /.
func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v]\t%s", r.Method, constants.HomeAPIRoute)
	fmt.Fprintf(w, "API seems fine.")
}

// GetAllPeople is handler for /text.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	grules.Engine.Execute(dataCtx, knowledgeBase)
	log.Printf("[%v]\t%s", r.Method, constants.GetAllPeopleAPIRoute)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&data.People); err != nil {
		log.Printf("[ERROR]\t%v", err)
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}
}

// UpdatePerson is handler for /person/:id.
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[%v]\t%s", r.Method, constants.UpdatePersonAPIRoute)
		http.Error(w, ":id should be an integer.", http.StatusBadRequest)
		return
	}

	log.Printf("[%v]\t%s\t:id = %v", r.Method, constants.UpdatePersonAPIRoute, id)

	var body models.UpdatePersonRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[ERROR]\t%v", err)
		http.Error(w, "Unable to parse json.", http.StatusBadRequest)
		return
	}

	for i, _ := range data.People {
		p := &data.People[i]

		if p.ID == int64(id) {
			p.Weight = body.Weight
			p.Height = body.Height
			log.Printf("[RUN]\tPerson with id `%v` updated.", id)
			fmt.Fprintf(w, "Person with id `%v` updated.", id)
			return
		}
	}
	http.Error(w, "Person with given id not found.", http.StatusNotFound)
}
