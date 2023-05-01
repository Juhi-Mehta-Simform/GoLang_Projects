package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   int
	Name string
	Book []Book
}

type Book struct {
	ID     int
	Title  string
	Author string
	UserID int
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Association sslmode=disable password=root port=5432")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Successful")
	}

	defer func() {
		db, err := db.DB()
		errCheck(err)
		db.Close()
		errCheck(err)
		fmt.Println("Closing database connection")
	}()

	db.AutoMigrate(&User{}, &Book{})
	//db.Create(&User{
	//	Name: "Juhi",
	//	Book: []Book{
	//		{
	//			Title:  "abcd",
	//			Author: "google",
	//		},
	//		{
	//			Title:  "ajdh",
	//			Author: "google",
	//		},
	//	},
	//})

	//err = db.Select("Name").Create(&User{}).Error
	//errCheck(err)

	//usr := User{
	//	Name: "abcd",
	//	Book: []Book{
	//		{
	//			Title:  "Golang",
	//			Author: "Google",
	//		},
	//	},
	//}

	//db.Omit("Book").Create(&usr)

	var user User
	//var books []Book
	db.Preload("Book").Find(&user, "id=1")
	//db.Model(&user).Association("Book").Find(&books)
	//fmt.Println(books)
	// err = db.Model(&user).Association("Book").Append(&Book{Title: "google", Author: "golang"})
	//fmt.Println(db.Model(&user).Where("user_id=1").Association("Book").Count())
	db.Model(&user).Association("Book").Clear()
	errCheck(err)
}
