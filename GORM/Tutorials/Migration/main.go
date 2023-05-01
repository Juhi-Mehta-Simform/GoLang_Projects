package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//type User struct {
//	ID      int
//	Name    string `gorm:"check:name_checker, name != 'sunny'"`
//	Address string
//}

type Person struct {
	gorm.Model
	CreditCard []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number   string
	PersonID int
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Migration sslmode=disable password=root port=5432")

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

	//db.AutoMigrate(&Person{}, &CreditCard{})
	err = db.Migrator().CreateTable(&Person{}, &CreditCard{})
	errCheck(err)
	//fmt.Println(db.Migrator().HasTable(&User{}))
	//fmt.Println(db.Migrator().HasTable("users"))
	//fmt.Println(db.Migrator().HasTable("names"))
	// db.Migrator().DropTable(&User{})
	// db.Migrator().RenameTable("users", "user_info")
	//db.Migrator().RenameTable("user_info", "users")
	// db.Migrator().AddColumn(&User{}, "Address")
	// db.Debug().Migrator().AlterColumn(&User{}, "Name")
	//db.Create(&User{
	//	Name: "sunny",
	//})
	//db.Migrator().CreateConstraint(&User{}, "name_checker")
	//db.Migrator().CreateConstraint(&Person{}, "CreditCard")
}
