package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type User struct {
	ID       int
	UserName string
	Order    []Order
}

type Order struct {
	ID     int
	UserID int
	Price  float64
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Preloading sslmode=disable password=root port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

	db.AutoMigrate(&User{}, &Order{})
	//db.Create(&User{
	//	UserName: "Juhi",
	//	Order: Order{
	//		Price: 100.00,
	//	},
	//})

	//db.Create(&User{
	//	UserName: "juhi",
	//	Order: []Order{
	//		{
	//			Price: 120.78,
	//		},
	//		{
	//			Price: 100.00,
	//		},
	//		{
	//			Price: 178.45,
	//		},
	//	},
	//})

	var users []User
	//db.Preload("Order").Find(&users)
	//err = db.Joins("Order").Find(&users).Error
	//err = db.Preload("Order", "user_id = ?", "1").Find(&users).Error
	//err = db.Where("user_name = ?", "Juhi").Preload("Order", "user_id = ?", "1").Find(&users).Error
	err = db.Preload("Order", func(db *gorm.DB) *gorm.DB {
		return db.Order("orders.price DESC")
	}).Find(&users).Error
	errCheck(err)
	fmt.Println(users)
}
