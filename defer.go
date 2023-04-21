package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
	disNum()
}

func disNum() {
	for i := 0; i <= 5; i++ {
		defer fmt.Print(i)
	}
}
