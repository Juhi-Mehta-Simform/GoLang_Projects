package main

import "fmt"

func main() {
	myNum := 12
	var ptr = &myNum
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(myNum)
}
