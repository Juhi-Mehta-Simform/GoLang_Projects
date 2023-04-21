package main

import "fmt"

func main() {
	var count int = 16
	if count > 10 {
		fmt.Println("Okay")
	} else if count < 10 {
		fmt.Println("Not okay")
	} else {
		fmt.Println("else")
	}
}
