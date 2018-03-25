package main

import (
	"fmt"
	"time"
)

var greetingsChannel = make(chan string)
var count = 1

func startChannel() {

	fmt.Printf("value=%v, type=%T\n", greetingsChannel, greetingsChannel)

	go listen()

	send()

}

func send() {
	for {
		greetingsChannel <- fmt.Sprintf("Hello %d", count)
		count++
		time.Sleep(2 * time.Second)
	}
}
func listen() {

	var greeting string
	for {
		fmt.Printf("Listening on channel...\n")
		greeting = <-greetingsChannel
		fmt.Printf("Received: value=%s, type=%T\n", greeting, greeting)
	}
}
