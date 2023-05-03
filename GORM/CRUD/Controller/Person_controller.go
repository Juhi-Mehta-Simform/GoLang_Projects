package Controller

import (
	"GORM/CRUD/Connection"
	Models2 "GORM/CRUD/Models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	var people []Models2.Person
	var books []Models2.Book
	fmt.Println("People")
	db := Connection.GetConnection()
	db.Find(&people)

	for i := range people {
		err := db.Model(&people[i]).Association("Books").Find(&books)
		errCheck(err)
		people[i].Books = books
	}
	fmt.Println(people)
	p, err := json.MarshalIndent(&people, "", "\t")
	errCheck(err)
	fmt.Fprint(writer, string(p))
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var (
		people []Models2.Person
		books  []Models2.Book
	)

	params := mux.Vars(request)
	fmt.Println("Person with Id: ", params["ID"])
	db := Connection.GetConnection()
	db.Find(&people, params["ID"])

	for i := range people {
		err := db.Model(&people[i]).Association("Books").Find(&books)
		errCheck(err)
		people[i].Books = books
	}
	fmt.Println(people)
	p, err := json.MarshalIndent(&people, "", "\t")
	errCheck(err)
	fmt.Fprintln(writer, "Person with Id: ", params["ID"])
	fmt.Fprint(writer, string(p))
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person Models2.Person
	db := Connection.GetConnection()
	err := json.NewDecoder(request.Body).Decode(&person)
	errCheck(err)
	db.Create(&person)
	err = json.NewEncoder(writer).Encode(&person)
	errCheck(err)
}

func UpdatePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var person Models2.Person
	db := Connection.GetConnection()
	db.First(&person, params["ID"])
	err := json.NewDecoder(request.Body).Decode(&person)
	errCheck(err)
	db.Save(&person)
	err = json.NewEncoder(writer).Encode(person)
	errCheck(err)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var person Models2.Person
	db := Connection.GetConnection()
	db.Delete(&person, params["ID"])
	err := json.NewEncoder(writer).Encode(&person)
	errCheck(err)
}
