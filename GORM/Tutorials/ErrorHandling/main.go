package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   int
	Name string `gorm:"unique"`
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=ErrorHandling sslmode=disable password=root port=5432")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Successful")
	}

	defer func() {
		db, err := db.DB()
		errCheck(err)
		err = db.Close()
		errCheck(err)
		fmt.Println("Closing database connection")
	}()

	err = db.AutoMigrate(&User{})
	errCheck(err)

	db.Create(&User{
		Name: "Juhi",
	})
}
