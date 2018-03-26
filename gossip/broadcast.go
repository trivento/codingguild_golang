package main

import (
	"log"
	"time"
)

// Make an array of max 4 with random nodes
func get4RandomMembers() []string {
	// TODO implement
	membersAsList := store.getMembersAsList()

	var gossipTo []string
	if len(membersAsList) < 4 {
		gossipTo = membersAsList
	} else {
		gossipTo = make([]string, 4)
		// TODO implement. Tip 'rand.Intn(4) kiest een random getal tussen 0 (inc) en 4 (excl)'
	}
	return gossipTo
}

// De broadcast daemon moet oneindig lang draaien en periodiek een gossip
// doen naar een aantal (max 4) members
func broadcast(commandChannel chan command) {
	log.Println("Starting the broadcast")
	for true {
		gossipTo := get4RandomMembers()
		log.Printf("Gossip to: %s", gossipTo)
		logKnownHosts()

		// Create the json with the structure: {"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
		// implemented in store.go
		m, _ := store.getMembersAsJSON()
		// gossip the known hosts to the gossiplist

		for _, node := range gossipTo {
			// Do not send to self
			// use http.Post
			// TODO implement
			log.Printf("send %s as json to %s", m, node)
		}
		time.Sleep(5 * time.Second)
	}
}

func logKnownHosts() {
	log.Printf("Known hosts:\n")
	for _, node := range store.getMembersAsList() {
		log.Printf("\t- %s\n", node)
	}
}
