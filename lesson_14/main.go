package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := make(map[string]string)

	map1["first_name"] = "Farhod"
	map1["surname"] = "Rajabov"
	map1["gender"] = "Male"

	jsoned, err := json.Marshal(map1) // [0->255]
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsoned))

	var person Person

	json.Unmarshal(jsoned, &person)
	fmt.Println(person.Firstname)
}

type Person struct {
	Firstname string `json:"first_name"`
	Surname   string `json:"surname"`
	Gender    string `json:"gender"`
}
