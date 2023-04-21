package main

import "fmt"

func main() {

	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
              
	for index, day := range days {
		fmt.Printf("index is %v & day is %v\n", index, day)
	}
}
