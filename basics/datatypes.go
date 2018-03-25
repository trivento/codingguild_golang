package main

import (
	"fmt"
)

func printDatatypes() {

	// const
	const (
		one = 1
		two = "2"
	)
	fmt.Printf("value=%d, type=%T\n", one, one)
	fmt.Printf("value=%s, type=%T\n", two, two)

	// array
	a := [3]int{1, 2, 3}
	fmt.Printf("value=%v, type=%T\n", a, a)

	// slice
	s := []int{1, 2, 3}
	fmt.Printf("value=%v, type=%T\n", s, s)
	s = append(s, 4)
	fmt.Printf("value=%v, type=%T\n", s, s)
	fmt.Printf("value=%v, type=%T\n", s[:2], s[:2])

	// map
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Printf("value=%v, type=%T\n", m, m)

	//struct
	type NaturalNumbers struct {
		one   string
		two   int
		three int
	}

	var countdown NaturalNumbers
	countdown.one = "1"
	countdown.two = 2
	countdown.three = 3
	fmt.Printf("value=%v, type=%T\n", countdown, countdown)

	countdownInWords := NaturalNumbers{
		one:   "one",
		two:   2,
		three: 3}
	fmt.Printf("value=%v, type=%T\n", countdownInWords, countdownInWords)

}
