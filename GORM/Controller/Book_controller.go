package Controller

import (
	"GORM/Connection"
	"GORM/Models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []Models.Book
	fmt.Println("Book")
	db := Connection.GetConnection()
	db.Find(&books)
	fmt.Println(books)
	p, err := json.MarshalIndent(&books, "", "\t")
	errCheck(err)
	fmt.Fprint(writer, string(p))
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	var books []Models.Book
	params := mux.Vars(request)
	fmt.Println("Book with ID: ", params["ID"])
	db := Connection.GetConnection()
	db.Find(&books, params["ID"])
	fmt.Println(books)

	p, err := json.MarshalIndent(&books, "", "\t")
	errCheck(err)
	fmt.Fprintln(writer, "Book with ID", params["ID"])
	fmt.Fprint(writer, string(p))
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book Models.Book
	db := Connection.GetConnection()
	err := json.NewDecoder(request.Body).Decode(&book)
	errCheck(err)
	db.Create(&book)
	err = json.NewEncoder(writer).Encode(&book)
	errCheck(err)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var book Models.Book
	db := Connection.GetConnection()
	db.First(&book, params["ID"])
	err := json.NewDecoder(request.Body).Decode(&book)
	errCheck(err)
	db.Save(&book)
	err = json.NewEncoder(writer).Encode(book)
	errCheck(err)
}

func DeleteBook(writer http.ResponseWriter, Request *http.Request) {
	params := mux.Vars(Request)
	var book Models.Book
	db := Connection.GetConnection()
	db.Delete(&book, params["ID"])
	err := json.NewEncoder(writer).Encode(&book)
	errCheck(err)
}
