package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID        int
	Name      string
	Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
	ID    int
	Name  string
	Users []User `gorm:"many2many:user_languages"`
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=manyToMany sslmode=disable password=root port=5432")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

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
		fmt.Println("Closing Database Connection")
	}()

	db.AutoMigrate(&User{}, &Language{})

	db.Create(&[]User{
		{
			Name: "Juhi",
			Languages: []Language{
				{Name: "Gujarati"},
				{Name: "Hindi"},
				{Name: "English"},
			},
		},
		{
			Name: "Janvi",
			Languages: []Language{
				{ID: 1, Name: "Gujarati"},
				{ID: 2, Name: "Hindi"},
				{Name: "Marathi"},
				{ID: 3, Name: "English"},
			},
		},
	})
}
