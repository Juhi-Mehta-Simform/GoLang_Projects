package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Name      []byte                 `gorm:"serializer:json"`
	Roles     Roles                  `gorm:"serializer:json"`
	Contracts map[string]interface{} `gorm:"serializer:json"`
	JobInfo   Job                    `gorm:"type:bytes;serializer:gob"`
}

type Roles []string

type Job struct {
	Title    string
	Location string
	IsIntern bool
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Serializer sslmode=disable password=root port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
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

	// db.AutoMigrate(&User{}, &Job{})
	//data := User{
	//	Name:      []byte("jinzhu"),
	//	Roles:     []string{"admin", "owner"},
	//	Contracts: map[string]interface{}{"name": "jinzhu", "age": 10},
	//	JobInfo: Job{
	//		Title:    "Developer",
	//		Location: "NY",
	//		IsIntern: false,
	//	},
	//}
	//
	//db.Create(data)
	var user []User
	db.Where(User{Name: []byte("jinzhu")}).Find(&user)

	for i := range user {
		fmt.Println(user[i])
	}

}
