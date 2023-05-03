package main

import "fmt"

const LoginId = 1200 //If The 1st character of variable is capital it shows it is public var.

func main() {
	var username string = "juhi"
	fmt.Println(username)
	fmt.Printf("Type of this variable: %T \n", username)

	var isName bool = true
	fmt.Println(isName)
	fmt.Printf("Type of this variable: %T \n", isName)

	var smallNum uint8 = 216
	fmt.Println(smallNum)
	fmt.Printf("Type of this variable: %T \n", smallNum)

	var smallFloat float32 = 45.12356383904567
	fmt.Println(smallFloat)
	fmt.Printf("Type of this variable: %T \n", smallFloat)

	numUser := 2000
	fmt.Println(numUser)
	fmt.Printf("Type of this variable: %T \n", numUser)

	fmt.Println(LoginId)
	fmt.Printf("Type of this variable: %T \n", LoginId)
}
