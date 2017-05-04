package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)


type Person struct {
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname  string `json:"lastname,omitempty"`
  Adress    *Adress `json:"adress,omitempty"`
}

type Adress struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
  Lastname string `json:"lastname,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func main() {
  router := mux.NewRouter()
  people = append(people, Person{ID: "1", Firstname: "john",  Lastname: "light", Adress: &Adress{City: "Dublin", State: "Carlifornia"} } )
  people = append(people, Person{ID: "2", Firstname: "maria",  Lastname: "hangman" } )
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
  log.Fatal(http.ListenAndServer(":12345", router))
}
