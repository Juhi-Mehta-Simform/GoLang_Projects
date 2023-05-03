package Models

import (
	"GORM/CRUD/Connection"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string
	Author   string
	CallNum  int `gorm:"unique_index"`
	PersonId int
}

func main() {
	var book Book
	db := Connection.GetConnection()
	db.Create(&book)
}
