package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Pick a valid number")
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if guess < 0 {
		fmt.Println("Please enter the positive number")
		return
	}
	// for turns := 0; turns < 5; turns++ {
	n := rand.Intn(guess + 1)
	if n == guess {
		fmt.Println("YOU WIN!!!")
		return
	}
	//}
	fmt.Println("YOU LOST...Try Again")
}
