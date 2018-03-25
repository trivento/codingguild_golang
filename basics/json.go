package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func printJSON() {

	//struct
	type Person struct {
		Name string
		Age  int
	}

	john := Person{"John", 21}
	lisa := Person{"Lisa", 22}

	fmt.Printf("value=%v, type=%T\n", john, john)
	fmt.Printf("value=%v, type=%T\n", lisa, lisa)

	johnJSON, _ := json.Marshal(john)
	fmt.Printf("value=%v, type=%T\n", johnJSON, johnJSON)
	fmt.Printf("value=%s, type=%T\n", johnJSON, johnJSON)

	var johnAgain Person
	_ = json.Unmarshal(johnJSON, &johnAgain)
	fmt.Printf("value=%v, type=%T\n", johnAgain, johnAgain)

	err := json.Unmarshal(johnJSON, &johnAgain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value=%v, type=%T\n", johnAgain, johnAgain)

}
