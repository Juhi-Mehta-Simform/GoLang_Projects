package main

import "fmt"

func main() {
	myUser := User{"Juhi", "juhi.m@gmail.com", 21}
	fmt.Println(myUser)
	myUser.GetAge()
	myUser.changeName()
}

type User struct {
	Name  string  
	Email string
	Age   int
}

func (u User) GetAge() {
	fmt.Println(u.Age)
}

func (u User) changeName() {
	u.Name = "Juhi Mehta"
	fmt.Println(u.Name)
}
