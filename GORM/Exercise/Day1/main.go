package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Employee struct {
	ID         int32      `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Number     int32      `json:"number,omitempty"`
	Employment Employment `gorm:"constraint:OnDelete:CASCADE" json:"employment"`
}

type Employment struct {
	ID         int32  `json:"id,omitempty"`
	EmployeeID int32  `json:"employee_id,omitempty"`
	Profile    string `json:"profile,omitempty"`
	Country    string `json:"country,omitempty"`
	JoinDate   string `gorm:"type:date" json:"join_date,omitempty"`
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Exercise sslmode=disable password=root port=5432")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		//PrepareStmt:            true,
		//SkipDefaultTransaction: true,
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

	// db.Migrator().DropTable(&Employee{}, &Employment{})
	//db.AutoMigrate(&Employee{}, &Employment{})
	//db.Create(&[]Employee{
	//	{
	//		ID:     101,
	//		Name:   "Juhi",
	//		Number: 642002,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "India",
	//			JoinDate: "2022/02/07",
	//		},
	//	},
	//	{
	//		ID:     102,
	//		Name:   "Janvi",
	//		Number: 130820,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "US",
	//			JoinDate: "2021/01/21",
	//		},
	//	},
	//	{
	//		ID:     103,
	//		Name:   "Dhruvi",
	//		Number: 151201,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "Germany",
	//			JoinDate: "2022/04/16",
	//		},
	//	},
	//	{
	//		ID:     104,
	//		Name:   "Keval",
	//		Number: 810001,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "India",
	//			JoinDate: "2021/11/19",
	//		},
	//	},
	//	{
	//		ID:     105,
	//		Name:   "Kishan",
	//		Number: 240202,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "US",
	//			JoinDate: "2022/08/07",
	//		},
	//	},
	//	{
	//		ID:     106,
	//		Name:   "Abhishek",
	//		Number: 240499,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "Thailand",
	//			JoinDate: "2021/05/27",
	//		},
	//	},
	//	{
	//		ID:     107,
	//		Name:   "Alana",
	//		Number: 342002,
	//		Employment: Employment{
	//			Profile:  "Software Engineer",
	//			Country:  "Germany",
	//			JoinDate: "2021/10/23",
	//		},
	//	},
	//	{
	//		ID:     108,
	//		Name:   "Sunny",
	//		Number: 160801,
	//		Employment: Employment{
	//			Profile:  "Sr. Software Engineer",
	//			Country:  "Australia",
	//			JoinDate: "2021/01/03",
	//		},
	//	},
	//})

	var emp []Employee
	//var country []Employment
	//db.Joins("Employment", db.Select("profile")).Where("country = ?", "India").Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	//db.Model(&Employment{}).Distinct("country").Order("country ASC").Find(&country).Limit(5)
	//for i := range country {
	//	fmt.Println(i+1, country[i].Country)
	//}

	//db.Joins("Employment").Where("country != ? and DATE_PART('Year', join_date) = ?", "India", 2022).Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	//db.Joins("Employment").Where("join_date BETWEEN ? AND ?", "2022-01-01", "2022-06-30").Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	//db.Joins("Employment").Where("country = ? and DATE_PART('Ye............................................................................................................ar', join_date) = ?", "Germany", 2021).Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	//db.Joins("Employment").Where("name LIKE ?", "__an%").Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	//db.Model(&Employee{}).Joins("Employment").Where("id = ? ", "101").Update("name", "Arjun")

	// db.Where("id = ?", "105").Delete(&Employee{})

	//db.Joins("Employment").Where("country IN ? ", []string{"India", "US"}).Order("id ASC").Find(&emp)
	//data, _ := json.MarshalIndent(emp, "", "\t")
	//fmt.Println(string(data))

	db.Joins("Employment", db.Select("profile", "country")).Find(&emp)
	data, _ := json.MarshalIndent(emp, "", "\t")
	fmt.Println(string(data))
}
