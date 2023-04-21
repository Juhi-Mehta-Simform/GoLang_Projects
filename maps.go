package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["Js"] = "JavaScript"
	languages["Py"] = "Python"
	languages["RoR"] = "Ruby on Rails"

	fmt.Println(languages)
	fmt.Println(languages["Js"])

	delete(languages, "Py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For the key %v, value is %v\n", key, value)
	}
}
