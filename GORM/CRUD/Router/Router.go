package Router

import (
	Controller2 "GORM/CRUD/Controller"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/people", Controller2.GetPeople).Methods("GET")
	Router.HandleFunc("/books", Controller2.GetBook).Methods("GET")
	Router.HandleFunc("/person/{ID}", Controller2.GetPerson).Methods("GET")
	Router.HandleFunc("/book/{ID}", Controller2.GetBook).Methods("GET")
	Router.HandleFunc("/createPerson", Controller2.CreatePerson).Methods("POST")
	Router.HandleFunc("/createBook", Controller2.CreateBook).Methods("POST")
	Router.HandleFunc("/updatePerson/{ID}", Controller2.UpdatePerson).Methods("PUT")
	Router.HandleFunc("/updateBook/{ID}", Controller2.UpdateBook).Methods("PUT")
	Router.HandleFunc("/deletePerson/{ID}", Controller2.DeletePerson).Methods("DELETE")
	Router.HandleFunc("/deleteBook/{ID}", Controller2.DeleteBook).Methods("DELETE")
}
