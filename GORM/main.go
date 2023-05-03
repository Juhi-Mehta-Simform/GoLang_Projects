package main

import (
	"GORM/CRUD/Connection"
	Models2 "GORM/CRUD/Models"
	"GORM/CRUD/Router"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	//person = &Models.Person{Name: "Juhi", Email: "juhi.m@simformsolutions.com"}
	//books  = []Models.Book{
	//	{Title: "Golang Basics", Author: "Google", CallNum: 1234, PersonId: 1},
	//	{Title: "Golang Advance", Author: "Google", CallNum: 5678, PersonId: 1},
	//}
	db  *gorm.DB
	err error
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := Connection.GetConnection()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection successful")
	}

	defer func() {
		db, err := db.DB()
		errCheck(err)
		err = db.Close()
		errCheck(err)
		fmt.Println("Closing database connection")
	}()

	err = db.AutoMigrate(&Models2.Person{}, &Models2.Book{})
	errCheck(err)

	log.Fatal(http.ListenAndServe(":8080", Router.Router))
}
