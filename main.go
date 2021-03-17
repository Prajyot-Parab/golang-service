package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage")
	fmt.Println("HomePage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", ReturnAllEmployees)
	myRouter.HandleFunc("/employee", CreateNewEmployee).Methods("POST")
	myRouter.HandleFunc("/employee/{id}", ReturnSingleEmployee).Methods("GET")
	myRouter.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	CheckConnection()
	handleRequests()
}
