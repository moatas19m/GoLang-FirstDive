package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Wizard struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	House  string  `json:"house"`
	Mentor *Mentor `json:"mentor"`
}

type Mentor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var wizards []Wizard

func main() {
	r := mux.NewRouter() // Exploring the gorilla mux library, r is now our Router

	wizards = append(wizards, Wizard{
		ID:     "1",
		Name:   "Harry Potter",
		House:  "Gryffindor",
		Mentor: &Mentor{FirstName: "Albus", LastName: "Dumbledore", Title: "The Headmaster"},
	})

	wizards = append(wizards, Wizard{
		ID:     "2",
		Name:   "Ron Weasley",
		House:  "Gryffindor",
		Mentor: &Mentor{FirstName: "Arther", LastName: "Weasley", Title: "The Father"},
	})

	r.HandleFunc("/wizards", getWizards).Methods("GET")
	r.HandleFunc("/wizards/{id}", getWizard).Methods("GET")
	r.HandleFunc("/wizards", createWizard).Methods("POST")
	r.HandleFunc("/wizards/{id}", updateWizard).Methods("PUT")
	r.HandleFunc("/wizards/{id}", deleteWizard).Methods("DELETE")

	fmt.Printf("Listening on port 9000\n")
	log.Fatal(http.ListenAndServe(":9000", r))

}

func deleteWizard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, wizard := range wizards {
		if wizard.ID == params["id"] {
			wizards = append(wizards[:idx], wizards[idx+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(wizards)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func updateWizard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, wizard := range wizards {
		if wizard.ID == params["id"] {
			wizards = append(wizards[:idx], wizards[idx+1:]...) // Delete the wizard with id passed in param, then create a new one
			var newWizard Wizard
			_ = json.NewDecoder(r.Body).Decode(&newWizard)
			newWizard.ID = params["id"]
			wizards = append(wizards, newWizard)
			err := json.NewEncoder(w).Encode(wizards)
			if err != nil {
				log.Fatal(err)
				return
			}
			return
		}
	}
}

func createWizard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newWizard Wizard
	_ = json.NewDecoder(r.Body).Decode(&newWizard)
	newWizard.ID = strconv.Itoa(rand.Intn(10000))
	wizards = append(wizards, newWizard)
	err := json.NewEncoder(w).Encode(wizards)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getWizard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, wizard := range wizards {
		if wizard.ID == params["id"] {
			err := json.NewEncoder(w).Encode(wizard)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func getWizards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(wizards)
	if err != nil {
		log.Fatal(err)
		return
	}
}
