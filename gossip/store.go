package main

import (
	"errors"
	"log"
)

// DELETE indicates delete command
var DELETE = "delete"

// ADD indicates update command
var ADD = "ADD"

// name can be delete, or update
type command struct {
	name       string
	memberlist []string
}

// De processLoop is een eeuwige lus die commands van het commandChan afhandeld
func processLoop(commandChan chan command) {
	for {
		command := <-commandChan
		log.Printf("Processing command %v", command)
		// TODO implement
	}
}

// Data structuur
//{"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
type members struct {
	Nodes []string `json:"nodes"`
}

// Store is de internal structure om network members in op te slaan

type Store struct {
	// maak een members attribuut. Tip: Go kent geen Set
	// TODO implement
}

// Maak een instantie van de store waarin alle members gezet worden.
var store = Store{
	// TODO implement
}

// voeg alle members in de memberlist toe aan de store
// Mag alleen aangeroepen worden vanuit de processLoop
func (store *Store) addMembers(memberlist []string) {
	// TODO implement
}

func (store *Store) getMembersAsList() []string {
	// TODO implement
	return nil
}

// Mag alleen aangeroepen worden vanuit de processLoop
func (store *Store) deleteMembers(memberlist []string) {
	// TODO implement
}

func (store *Store) getMembersAsJSON() ([]byte, error) {
	// TODO implement
	return nil, errors.New("not implemented")
}
