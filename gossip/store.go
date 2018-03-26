package main

import (
	"encoding/json"
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

func processLoop(commandChan chan command) {
	for {
		command := <-commandChan
		log.Printf("Processing command %v", command)
		if command.name == DELETE {
			store.deleteMembers(command.memberlist)
		} else if command.name == ADD {
			store.addMembers(command.memberlist)
		}

	}
}

// Data structuur
//{"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
type members struct {
	Nodes []string `json:"nodes"`
}

// Store is the internal structure to store all network members
type Store struct {
	members map[string]bool
}

// Maak een store waarin alle members gezet worden. Tip: Go kent geen Set
var store = Store{make(map[string]bool)}

func (store *Store) addMembers(memberlist []string) {
	for _, m := range memberlist {
		store.members[m] = true
	}
}

func (store *Store) getMembersAsList() []string {
	membersAsList := make([]string, len(store.members))
	idx := 0
	for node := range store.members {
		membersAsList[idx] = node
		idx++
	}
	return membersAsList
}

func (store *Store) deleteMembers(memberlist []string) {
	for _, m := range memberlist {
		delete(store.members, m)
	}
}

func (store *Store) getMembersAsJSON() ([]byte, error) {
	return json.Marshal(members{store.getMembersAsList()})
}
