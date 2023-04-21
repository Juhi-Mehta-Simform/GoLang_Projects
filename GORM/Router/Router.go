package Router

import (
	"GORM/Controller"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/people", Controller.GetPeople).Methods("GET")
	Router.HandleFunc("/books", Controller.GetBook).Methods("GET")
	Router.HandleFunc("/person/{ID}", Controller.GetPerson).Methods("GET")
	Router.HandleFunc("/book/{ID}", Controller.GetBook).Methods("GET")
	Router.HandleFunc("/createPerson", Controller.CreatePerson).Methods("POST")
	Router.HandleFunc("/createBook", Controller.CreateBook).Methods("POST")
	Router.HandleFunc("/updatePerson/{ID}", Controller.UpdatePerson).Methods("PUT")
	Router.HandleFunc("/updateBook/{ID}", Controller.UpdateBook).Methods("PUT")
	Router.HandleFunc("/deletePerson/{ID}", Controller.DeletePerson).Methods("DELETE")
	Router.HandleFunc("/deleteBook/{ID}", Controller.DeleteBook).Methods("DELETE")
}
