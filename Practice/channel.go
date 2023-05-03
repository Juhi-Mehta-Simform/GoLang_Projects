package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

	// myChannel <- 5
	// fmt.Println(<-myChannel)
}
