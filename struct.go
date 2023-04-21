package main

import "fmt"

func main() {
	myUser := User{"Juhi", "juhi@gmail.com", true, 21}
	fmt.Println(myUser)
	fmt.Println(myUser.Name)
	fmt.Printf("The details are: %+v\n", myUser)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
