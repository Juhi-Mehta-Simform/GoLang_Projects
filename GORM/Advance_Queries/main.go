package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"type-varchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

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
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=GORM sslmode=disable password=root port=5432")
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

	err = db.AutoMigrate(&Person{}, &Book{})
	errCheck(err)

	//var person = &Person{Name: "Juhi", Email: "juhi@gmail.com"}
	//db.Create(&person)

	//person := []*Person{
	//	{Name: "janvi", Email: "janvi@gmail"},
	//	{Name: "dhruvi", Email: "dhruvi@gmail"},
	//}
	//db.Create(&person)

	// var person []Person
	//db.Take(&person)
	//db.Select("Name").Find(&person)
	//db.Last(&person)
	//for i := range person {
	//	fmt.Println(person[i])
	//}
	//db.Where("name = ?", "Juhi").Find(&person)
	//db.Find(&person, "id = 1")
	// fmt.Println(person)

	//books := []Book{
	//	{Title: "abcd", Author: "hjgj", PersonID: 1},
	//	{Title: "dbca", Author: "hjgj", PersonID: 1},
	//	{Title: "cdab", Author: "hjgj", PersonID: 2},
	//	{Title: "bcde", Author: "hjgj", PersonID: 3},
	//}
	//db.Create(&books)

	//var books []Book
	// db.Where("person_id = 1").Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Find(&books)
	// db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Limit(1).Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Offset(2).Find(&books)
	//for _, book := range books {
	//	fmt.Println(book)
	//}

	var result []int
	db.Model(&Book{}).Select("count(*)").Group("person_id").Find(&result)
	fmt.Println(result)
}
