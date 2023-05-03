package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://lco.dev",
		"https://facebook.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		getStatusCode(web)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d statuscode for %s\n", res.StatusCode, endpoint)
	}
}
