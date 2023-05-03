package main

import "fmt"

func main() {
	fmt.Println("Enter the number: ")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
