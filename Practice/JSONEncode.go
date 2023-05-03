package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"juhi", "1234", nil},
		{"janvi", "abcd", permissions{"admin": true}},
		{"dhruvi", "ab34", permissions{"write": true}},
	}

	// out, err := json.Marshal(users)
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
