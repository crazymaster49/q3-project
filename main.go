package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}
func DeletePeopleEndPoint(w http.ResponseWriter, req *http.Request) {

}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Nic", LastName: "Andra", Address: &Address{City: "Parker", State: "CO"}})
	people = append(people, Person{ID: "2", FirstName: "Dan", LastName: "Prodo", Address: &Address{City: "Parker", State: "CO"}})
	people = append(people, Person{ID: "3", FirstName: "Sam", LastName: "Conner"})
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1738", router))
}
