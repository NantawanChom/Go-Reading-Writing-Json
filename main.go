package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	my_json := `
		[
			{
				"first_name": "John",
				"last_name": "Kent",
				"hair_color": "Gray",
				"has_dog": false
			},
			{
				"first_name": "Mary",
				"last_name": "Wayne",
				"hair_color": "Red",
				"has_dog": true
			}
		]
	`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(my_json), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from  struct

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Bob"
	m1.LastName = "Hello"
	m1.HairColor = "black"
	m1.HasDog = true

	var m2 Person

	m2.FirstName = "Anna"
	m2.LastName = "Wile"
	m2.HairColor = "Black"
	m2.HasDog = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
