package main

import "fmt"

func main() {
	a := 0
	b := 1
	next_term := 0
	var n int

	fmt.Printf("Enter the total number of terms: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Println(" ", a)
			continue
		}
		if i == 2 {
			fmt.Println(" ", b)
			continue
		}
		next_term = a + b
		a = b
		b = next_term
		fmt.Println(" ", next_term)
	}
}
