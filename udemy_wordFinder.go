package main

import (
	"fmt"
	"strings"
)

const sentence = "lazy cat jumps again and again and again"

// func main() {
// 	words := strings.Fields(sentence)
// 	query := os.Args[1:]
// queries:
// 	for _, q := range query {
// 		for i, w := range words {
// 			if q == w {
// 				fmt.Printf("%-2d: %q\n", i+1, w)
// 				break queries
// 			}
// 		}
// 	}
// }

func main() {
	var query string
	fmt.Scanln(&query)
	query = strings.ToLower(query)
	queries := strings.Split(query, ",")
	words := strings.Fields(sentence)

	for _, q := range queries {
		index := 0
		for i, w := range words {
			if q == w {
				index = i + 1
				fmt.Printf("%-2d: %q\n", index, w)
				break
			}
		}
		if index == 0 {
			fmt.Printf("%q is not found\n", q)
			continue
		}
	}
}
