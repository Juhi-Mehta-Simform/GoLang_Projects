package main

import "fmt"

//
//func main() {
//	myUser := User{"Juhi", "juhi@gmail.com", true, 21}
//	fmt.Println(myUser)
//	fmt.Println(myUser.Name)
//	fmt.Printf("The details are: %+v\n", myUser)
//}
//
//type User struct {
//	Name   string
//	Email  string
//	Status bool
//	Age    int
//}

type person struct {
	name string
	age  int
}

func main() {
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
