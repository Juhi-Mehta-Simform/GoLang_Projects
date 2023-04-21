package main

import (
	"fmt"
	"sort"
)

func main() {
	var arrayList = []string{"apple", "banana", "peach"}
	fmt.Println(arrayList)
	fmt.Printf("type: %T\n", arrayList)
	arrayList = append(arrayList, "Mango", "strawberry")
	fmt.Println(arrayList)
	fmt.Println(arrayList[1:])
	fmt.Println(arrayList[1:3])
	fmt.Println(arrayList[:3])

	var numArray = make([]int, 4)
	numArray[0] = 123
	numArray[1] = 648
	numArray[2] = 453
	numArray[3] = 321

	fmt.Println(numArray)
	numArray = append(numArray, 555, 666, 777)
	fmt.Println(numArray)
	sort.Ints(numArray)
	fmt.Println(numArray)
	fmt.Println(sort.IntsAreSorted(numArray))

	var index int = 2
	numArray = append(numArray[:index], numArray[index+1:]...)
	fmt.Println(numArray)
}
