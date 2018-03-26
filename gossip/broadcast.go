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

// De broadcast daemon moet oneindig lang draaien en periodiek een gossip
// doen naar een aantal (max 4) members
func broadcast() {
	log.Println("Starting the broadcast")
	for true {
		m, _ := getMembers()
		logKnownHosts()
		var gossipTo []string

		membersAsList := make([]string, len(store))
		idx := 0
		for node := range store {
			membersAsList[idx] = node
			idx++
		}

		if len(membersAsList) < 4 {
			gossipTo = membersAsList
		} else {
			gossipTo = make([]string, 4)
			for idx := range gossipTo {
				pick := rand.Intn(len(membersAsList))
				gossipTo[idx] = membersAsList[pick]
			}
		}

		for _, node := range gossipTo {
			// Do not send to self
			if !strings.HasPrefix(node, myHost) {
				logline := fmt.Sprintf("Sending to %s", node)
				r, e := http.Post(node+"/members", "application/json", bytes.NewReader(m))
				if e != nil {
					logline = fmt.Sprintf("%s. Error: %s", logline, e.Error())
					delete(store, node)
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
	for node := range store {
		log.Printf("\t- %s\n", node)
	}
}
