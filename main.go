package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"fristname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEnpoint(w http.ResponseWriter, req *http.Request) {
	paras := mux.Vars(req)
	for _, item := range people {
		if item.ID == paras["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(Person{})
}

func CreatePersonEnpoint(w http.ResponseWriter, req *http.Request) {
	paras := mux.Vars(req)
	var per Person
	json.NewDecoder(req.Body).Decode(&per)
	per.ID = paras["id"]
	people = append(people, per)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEnpoint(w http.ResponseWriter, req *http.Request) {
	paras := mux.Vars(req)
	for k, item := range people {
		if item.ID == paras["id"] {
			people = append(people[:k], people[k+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func GetPeopleEnpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Println("test")
}

func main() {
	r := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Jeffrey", Lastname: "Li", Address: &Address{City: "Chengdu", State: "Sichuan"}})
	people = append(people, Person{ID: "2", Firstname: "Aoxuan", Lastname: "Li", Address: &Address{City: "Chengdu", State: "Sichuan"}})

	r.HandleFunc("/people/{id}", GetPersonEnpoint).Methods("GET")
	r.HandleFunc("/people/{id}", CreatePersonEnpoint).Methods("POST")
	r.HandleFunc("/people/{id}", DeletePersonEnpoint).Methods("DELETE")
	r.HandleFunc("/people", GetPeopleEnpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", r))
}
