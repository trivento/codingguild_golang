package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Make an array of max 4 with random nodes
func get4RandomMembers() []string {
	membersAsList := store.getMembersAsList()

	var gossipTo []string
	if len(membersAsList) < 4 {
		gossipTo = membersAsList
	} else {
		gossipTo = make([]string, 4)
		for idx := range gossipTo {
			pick := rand.Intn(len(membersAsList))
			gossipTo[idx] = membersAsList[pick]
		}
	}
	return gossipTo
}

// De broadcast daemon moet oneindig lang draaien en periodiek een gossip
// doen naar een aantal (max 4) members
func broadcast(commandChannel chan command) {
	log.Println("Starting the broadcast")
	for true {
		m, _ := store.getMembersAsJSON()
		logKnownHosts()

		gossipTo := get4RandomMembers()

		for _, node := range gossipTo {
			// Do not send to self
			if !strings.HasPrefix(node, myHost) {
				logline := fmt.Sprintf("Sending to %s", node)
				r, e := http.Post(node+"/members", "application/json", bytes.NewReader(m))
				if e != nil {
					logline = fmt.Sprintf("%s. Error: %s", logline, e.Error())
					commandChannel <- command{DELETE, []string{node}}
				} else {
					logline = fmt.Sprintf("%s. Result: [%s]", logline, r.Status)
				}
				log.Print(logline)
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func logKnownHosts() {
	log.Printf("Known hosts:\n")
	for node := range store.members {
		log.Printf("\t- %s\n", node)
	}
}
