package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//type User struct {
//	gorm.Model
//	Name       string     `gorm:"unique"`
//	CreditCard CreditCard `gorm:"foreignKey:UserName;references:name"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	UserName string
//}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=hasOne sslmode=disable password=root port=5432")

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
		err = db.Close()
		errCheck(err)
		fmt.Println("Closing Database Connection")
	}()

	//db.AutoMigrate(&User{}, &CreditCard{})
	db.AutoMigrate(&Toy{}, &Dog{}, &Cat{})
}
