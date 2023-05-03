package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var max int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&max)
	var uniques []int
	// const max = 5
	// var uniques [max]int
loop:
	// for found := 0; found < max; {
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		// fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}
		// uniques[found] = n
		uniques = append(uniques, n)
		// found++
	}
	fmt.Println("uniques: ", uniques)

}
