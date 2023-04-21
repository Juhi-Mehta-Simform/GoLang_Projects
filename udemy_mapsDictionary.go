package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemme",
	}
	dict["up"] = "yukarl"
	dict["down"] = "asagl"
	dict["mistake"] = ""

	value, ok := dict[query]
	if !ok {
		fmt.Printf("%q not found\n", query)
		return
	}
	fmt.Printf("%q means %#v\n", query, value)
	// english := []string{"good", "great", "perfect"}
	// turkish := []string{"iyi", "harika", "mukemmel"}

	// for i, w := range english {
	// 	if query == w {
	// 		fmt.Printf("%q means %q\n", w, turkish[i])
	// 		return
	// 	}
	// }
}
