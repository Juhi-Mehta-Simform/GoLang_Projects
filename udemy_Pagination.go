package main

import "fmt"

func main() {
	items := []string{
		"abc", "bcd", "cde", "def",
		"efg", "fgh", "ghi", "hij",
		"ijk", "jkl", "klm", "lmn",
		"mno", "nop",
	}
	// fmt.Println("0:4\n", items[0:4])
	// fmt.Println("4:8\n", items[4:8])
	// fmt.Println("8:12\n", items[8:12])
	// fmt.Println("12:13\n", items[12:13])
	const pageSize = 4
	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]
		fmt.Printf("page #%d", (from / pageSize))
		fmt.Println(currentPage)
	}
}
