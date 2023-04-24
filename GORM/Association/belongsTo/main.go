package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	UserID    int `gorm:"primaryKey"`
	Name      string
	CompanyID int
	Company   Company `gorm:"references:Number;foreignKey:CompanyID"`
}

type Company struct {
	Number int `gorm:"primaryKey"`
	Name   string
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=posgres dbname=BelongsTo sslmode=disable password=root port=5432")
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
	err = db.AutoMigrate(&Company{}, &User{})
	errCheck(err)
	//db.Create(&User{
	//	Name:    "Juhi",
	//	Company: Company{Name: "Simform", Number: 5},
	//})
	// db.Where("user_id=2").Delete(&User{})
	// db.Where("number=5").Delete(&Company{})

	var getUser []User
	db.Joins("Company").Find(&getUser)
	for i := range getUser {
		fmt.Println(getUser[i])
	}
}
