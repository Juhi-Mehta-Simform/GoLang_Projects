package main

import "fmt"

func main() {
	var arrayList [4]string
	arrayList[0] = "apple"
	arrayList[1] = "banana"
	arrayList[3] = "peach"
	fmt.Println(arrayList)
	fmt.Println(len(arrayList))

	var vegList = [3]string{"patoto", "beans", "tomato"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
