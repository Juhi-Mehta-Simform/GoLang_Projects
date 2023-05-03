package Models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string
	Author   string
	CallNum  int `gorm:"unique_index"`
	PersonId int
}
