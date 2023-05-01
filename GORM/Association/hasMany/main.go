package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID         int
	Name       string
	CreditCard []CreditCard
}

type CreditCard struct {
	ID     int
	Number string
	UserID int
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=hasMany sslmode=disable password=root port=5432")
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

	db.AutoMigrate(&User{}, &CreditCard{})

	//db.Create(&[]User{
	//	{
	//		Name: "Juhi",
	//		CreditCard: []CreditCard{
	//			{
	//				Number: "112233445566",
	//			},
	//			{
	//				Number: "111122223333",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Janvi",
	//		CreditCard: []CreditCard{
	//			{
	//				Number: "222233334444",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Dhruvi",
	//		CreditCard: []CreditCard{
	//			{
	//				Number: "333344445555",
	//			},
	//			{
	//				Number: "444455556666",
	//			},
	//			{
	//				Number: "555566667777",
	//			},
	//		},
	//	},
	//})

	//err = db.Where("id>=4").Delete(&User{}).Error
	//errCheck(err)
	var users []User
	err = db.Model(&User{}).Preload("CreditCard").Find(&users).Error
	errCheck(err)
	for i := range users {
		fmt.Println(users[i])
	}
}
