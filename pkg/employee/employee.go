package employee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host string = os.Getenv("DB_URL")
var port int = 5432
var user string = os.Getenv("DB_USER")
var password string = os.Getenv("DB_PASSWORD")
var dbname string = os.Getenv("DB_NAME")

type Employee struct {
	gorm.Model
	Id   string
	Name string
	Role string
	Org  string
}

func CheckConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to Database")
	}
	_ = db
	fmt.Println("Connected")
}

func ReturnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ReturnAllEmployees endpoint hit")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	var employees []Employee
	db.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func ReturnSingleEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ReturnSingleEmployee endpoint hit")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	vars := mux.Vars(r)
	key := vars["id"]
	var employee Employee

	db.Where("id = ?", key).Find(&employee)
	json.NewEncoder(w).Encode(employee)
}

func CreateNewEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createNewEmployee endpoint hit")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to Database")
	}

	var emp Employee
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &emp)

	db.Create(&Employee{Id: emp.Id, Name: emp.Name, Role: emp.Role, Org: emp.Org})
	fmt.Fprintf(w, "User succesfully created")
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteEmployee endpoint hit")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var emp Employee
	db.Where("id = ?", id).Find(&emp)
	db.Delete(&emp)
	fmt.Fprintf(w, "User succesfully deleted")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEmployee endpoint hit")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	vars := mux.Vars(r)
	key := vars["id"]

	var emp, empData Employee
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &emp)

	db.Where("id = ?", key).Find(&empData)

	empData.Name = emp.Name
	empData.Role = emp.Role
	empData.Org = emp.Org

	db.Save(&empData)
	fmt.Fprintf(w, "User succesfully updated")
}
