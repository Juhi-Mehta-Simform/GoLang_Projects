package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()
	fmt.Println(currentDate)
	fmt.Println(currentDate.Format("01-02-2006 Monday 15:04:05")) //MM-DD-YYYY
	fmt.Println(currentDate.Format("02-01-2006 Monday 15:04:05")) //DD-MM-YYYY
	createdDate := time.Date(2021, time.April, 6, 8, 58, 15, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday 15:04:05"))
}
