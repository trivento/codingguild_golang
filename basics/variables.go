package main

import (
	"fmt"
)

func printVariables() {

	// declare and assign
	one := "1"
	two := 2

	// declare, and assign
	var three = "3"
	var four int
	four = 4

	fmt.Printf("value=%s, type=%T\n", one, one)
	fmt.Printf("value=%d, type=%T\n", two, two)
	fmt.Printf("value=%s, type=%T\n", three, three)
	fmt.Printf("value=%d, type=%T\n", four, four)

}
