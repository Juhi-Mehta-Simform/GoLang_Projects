package main

import "fmt"

func main() {
	greet()
	result := addition(3, 7)
	fmt.Println("Result is: ", result)
	addNum := sumNum(4, 5, 3)
	fmt.Println(addNum)
	sumResult, myMsg := sum(2, 6, 8, 1, 5, 3)
	fmt.Println(sumResult)
	fmt.Println(myMsg)
}

func greet() {
	fmt.Println("Hello, Good afternoon")
}

func addition(valOne int, valTwo int) int {
	return valOne + valTwo
}

func sumNum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func sum(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hello"
}
