package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string `json: "name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"last_name"`
	age        int    `json:"age"`
}

func main() {
	jsonUser := `{"name": "John", "last_name": "Doe", "middle_name": "M", "age": 42}`
	var u User
	err := json.Unmarshal([]byte(jsonUser), &u)
	fmt.Println(err)
	fmt.Println(u)

	res, _ := json.Marshal(u)
	fmt.Println(string(res))
}
