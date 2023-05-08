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
	DeptName   string     `gorm:"type:varchar(50)" json:"dept_name,omitempty"`
	Salary     int32      `json:"salary,omitempty"`
	Department Department `gorm:"references:EmpDept;foreignKey:DeptName;constraint:OnDelete:CASCADE, onUpdate:CASCADE" json:"department"`
}

type Department struct {
	EmpDept  string `gorm:"primaryKey" json:"emp_dept,omitempty"`
	Location string `json:"location,omitempty"`
}

type EmployeeIndia struct {
	ID         int32
	Name       string
	Gender     string
	Department string
}

type EmployeeUK struct {
	ID         int32
	Name       string
	Gender     string
	Department string
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Exercise2 sslmode=disable password=root port=5432")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		//Logger:                 logger.Default.LogMode(logger.Info),
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

	db.AutoMigrate(&Department{}, &Employee{})
	db.AutoMigrate(&EmployeeIndia{}, &EmployeeUK{})
	//db.Create(&[]Employee{
	//	{
	//		ID:       101,
	//		Name:     "Mohan",
	//		DeptName: "Admin",
	//		Salary:   4000,
	//		Department: Department{
	//			EmpDept:  "Admin",
	//			Location: "Mumbai",
	//		},
	//	},
	//	{
	//		ID:       102,
	//		Name:     "Rajkumar",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       103,
	//		Name:     "Akbar",
	//		DeptName: "IT",
	//		Salary:   4000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       104,
	//		Name:     "Dorvin",
	//		DeptName: "Finance",
	//		Salary:   6500,
	//		Department: Department{
	//			EmpDept:  "Finance",
	//			Location: "Ahmedabad",
	//		},
	//	},
	//	{
	//		ID:       105,
	//		Name:     "Rohit",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       106,
	//		Name:     "Rajesh",
	//		DeptName: "Finance",
	//		Salary:   5000,
	//		Department: Department{
	//			EmpDept:  "Finance",
	//			Location: "Ahmedabad",
	//		},
	//	},
	//	{
	//		ID:       107,
	//		Name:     "Preet",
	//		DeptName: "HR",
	//		Salary:   7000,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       108,
	//		Name:     "Maryam",
	//		DeptName: "Admin",
	//		Salary:   4000,
	//		Department: Department{
	//			EmpDept:  "Admin",
	//			Location: "Mumbai",
	//		},
	//	},
	//	{
	//		ID:       109,
	//		Name:     "Sanjay",
	//		DeptName: "IT",
	//		Salary:   6500,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       110,
	//		Name:     "Vasudha",
	//		DeptName: "IT",
	//		Salary:   7000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       111,
	//		Name:     "Melinda",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       112,
	//		Name:     "Komal",
	//		DeptName: "IT",
	//		Salary:   10000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       113,
	//		Name:     "Gautham",
	//		DeptName: "Admin",
	//		Salary:   2000,
	//		Department: Department{
	//			EmpDept:  "Admin",
	//			Location: "Mumbai",
	//		},
	//	},
	//	{
	//		ID:       114,
	//		Name:     "Manisha",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       115,
	//		Name:     "Chandni",
	//		DeptName: "IT",
	//		Salary:   4500,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       116,
	//		Name:     "Satya",
	//		DeptName: "Finance",
	//		Salary:   6500,
	//		Department: Department{
	//			EmpDept:  "Finance",
	//			Location: "Ahmedabad",
	//		},
	//	},
	//	{
	//		ID:       117,
	//		Name:     "Adarsh",
	//		DeptName: "HR",
	//		Salary:   3500,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       118,
	//		Name:     "Tejaswi",
	//		DeptName: "Finance",
	//		Salary:   5500,
	//		Department: Department{
	//			EmpDept:  "Finance",
	//			Location: "Ahmedabad",
	//		},
	//	},
	//	{
	//		ID:       119,
	//		Name:     "Cory",
	//		DeptName: "HR",
	//		Salary:   8000,
	//		Department: Department{
	//			EmpDept:  "HR",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       120,
	//		Name:     "Monica",
	//		DeptName: "Admin",
	//		Salary:   5000,
	//		Department: Department{
	//			EmpDept:  "Admin",
	//			Location: "Mumbai",
	//		},
	//	},
	//	{
	//		ID:       121,
	//		Name:     "Rosalin",
	//		DeptName: "IT",
	//		Salary:   6000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       122,
	//		Name:     "Ibrahim",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       123,
	//		Name:     "Vikram",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//	{
	//		ID:       124,
	//		Name:     "Dheeraj",
	//		DeptName: "IT",
	//		Salary:   11000,
	//		Department: Department{
	//			EmpDept:  "IT",
	//			Location: "Banglore",
	//		},
	//	},
	//})
	//db.Create(&Department{
	//	EmpDept:  "Marketing",
	//	Location: "Pune",
	//})

	//db.Create([]EmployeeIndia{
	//	{
	//		Name:       "Pranaya",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Priyanka",
	//		Gender:     "Female",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Preety",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Subrat",
	//		Gender:     "Male",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Anurag",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//})
	//
	//db.Create([]EmployeeUK{
	//	{
	//		Name:       "James",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Priyanka",
	//		Gender:     "Female",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Sara",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Subrat",
	//		Gender:     "Male",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Pam",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//})

	var emp []Employee
	//sal := []map[string]interface{}{}
	//var count int64

	//db.Model(&Employee{}).InnerJoins("Department", db.Where("dept_name = ? and location = ?", "HR", "Banglore")).Select("MAX(salary) as max, MIN(salary) as min").Group("emp_dept").Find(&sal)

	//db.Model(&Employee{}).InnerJoins("Department", db.Where("dept_name = ? OR dept_name = ?", "IT", "Finance")).Count(&count)

	//db.Model(&Employee{}).InnerJoins("Department").Select("AVG(salary) as average").Group("emp_dept").Find(&sal)
	//data, _ := json.MarshalIndent(sal, "", "\t")
	//fmt.Println(string(data))

	//db.Model(&Employee{}).InnerJoins("Department", db.Select("location")).Select("SUM(salary) as sum").Group("location").Find(&sal)

	//db.Model(&Employee{}).Select("dept_name, COUNT(*)").Group("dept_name").Having("COUNT(id) > ?", "4").Find(&sal)

	q1 := db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.emp_dept").Select("MAX(salary)").Where("dp.location=?", "Banglore")
	q2 := db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.emp_dept").Select("MIN(salary)").Where("dp.location=?", "Mumbai")
	db.Preload("Department").Where("salary IN ((?), (?))", q1, q2).Find(&emp)

	//db.Raw("Create table employee_hr_admin AS SELECT id, name, dept_name, location FROM employees inner join departments on dept_name=emp_dept WHERE emp_dept IN ('HR','Admin')").Row()

	//db.Raw("insert into employee_hr_admin ?", db.Model(&Employee{}).Joins("inner join departments on dept_name=emp_dept").Select("id, name, emp_dept, location").Where("emp_dept = ?", "IT")).Row()

	//db.Model(&Employee{}).Joins("Department").Where("salary > (?)",
	//	db.Model(Employee{}).Select("AVG(salary)").Where("dept_name = emp_dept").Group("emp_dept"),
	//).Find(&emp)

	//db.Model(Employee{}).Joins("Department").Where("salary > (?)",
	//	db.Model(Employee{}).Select("salary").Where("id = 109"),
	//).Find(&emp)

	//db.Model(Employee{}).Select("name, dept_name, case when salary > (@avg) then 'Higher' when salary < (@avg) then 'Lower' when salary = (@avg) then 'Equal' END as relative_salary", sql.Named("avg", db.Select("avg(salary)").Table("employees"))).Find(&sal)

	//db.Raw("(?) union (?)", db.Model(&EmployeeIndia{}).Select("employee_india.*, 'India' as country_name"), db.Model(&EmployeeUK{}).Select("employee_uks.*, 'UK' as country_name")).Find(&sal)

	//db.Model(&EmployeeUK{}).Not("name IN (?)", db.Model(&EmployeeIndia{}).Select("name")).Find(&sal)

	//db.Model(&EmployeeUK{}).Where("name IN (?)", db.Model(&EmployeeIndia{}).Select("name")).Find(&sal)
	data, _ := json.MarshalIndent(emp, "", "\t")
	fmt.Println(string(data))
}
