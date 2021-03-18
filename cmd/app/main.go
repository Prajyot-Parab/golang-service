package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Prajyot-Parab/golang-service/pkg/employee"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage")
	fmt.Println("HomePage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", employee.ReturnAllEmployees)
	myRouter.HandleFunc("/employee", employee.CreateNewEmployee).Methods("POST")
	myRouter.HandleFunc("/employee/{id}", employee.ReturnSingleEmployee).Methods("GET")
	myRouter.HandleFunc("/employee/{id}", employee.DeleteEmployee).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}", employee.UpdateEmployee).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	file, err := os.OpenFile("./logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Checking Connection...")
	employee.CheckConnection()
	handleRequests()
	log.Println("Application Running...")
}
