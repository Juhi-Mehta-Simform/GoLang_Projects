package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	ID   int
	Name string
	Book []Book
}

type Book struct {
	ID       int
	Title    string
	Author   string
	PersonID int
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Hooks sslmode=disable password=root port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to Database")
	}

	defer func() {
		db, err := db.DB()
		errCheck(err)
		err = db.Close()
		errCheck(err)
		fmt.Println("Closing Database Connection")
	}()

	err = db.AutoMigrate(&Person{}, &Book{})
	errCheck(err)

	//var person Person
	//person = Person{
	//	Name: "Janvi",
	//	Book: []Book{
	//		{
	//			Title:  "golang",
	//			Author: "Google",
	//		},
	//	},
	//}
	CreateUser(db)
	//var user Person
	//db.Find(&user)
	//fmt.Println(user)
}

//func (p *Person) BeforeCreate(tx *gorm.DB) error {
//	fmt.Println(*p)
//	return nil
//}
//
//func (p *Person) BeforeSave(tx *gorm.DB) error {
//	fmt.Println(*p)
//	return nil
//}
//
//func (p *Person) AfterCreate(db *gorm.DB) error {
//	var person Person
//	db.Model(p).Preload("Book").Find(&person)
//	fmt.Println(person)
//	return nil
//}

func CreateUser(db *gorm.DB) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&Person{Name: "abcd", Book: []Book{{Title: "jsfg", Author: "google"}}}).Error
	errCheck(err)
	tx.SavePoint("abcd")
	err = tx.Create(&Person{Name: "dcba"}).Error
	errCheck(err)
	tx.RollbackTo("abcd")
	tx.Commit()
}
