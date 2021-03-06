package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/trivento/codingguild_golang/iptools"
)

/*
Opdracht:

Implementeer een netwerk cluster member die zich kan aanmelden bij een cluster van Nodes en
actief (gossippen van kennis over het cluster) participeert.

- Gossip protocol:
	1. Meld je aan bij de Seed node door een lijst van known nodes met enkel jezelf te sturen
	2. Periodiek naar een aantal nodes in de lijst van members jouw kennis van het netwerk sturen
	3. Als je een gossip/lijst van nodes ontvangt, dan voeg je alle nieuwe nodes toe aan je eigen cluster kennis
- Data structuur van de informatie die ge-gossiped wordt:
	vb: {"nodes":["http://10.248.30.150:8082","http://10.248.30.143:8081"]}


*/

var myHost string

func handleGossip(w http.ResponseWriter, r *http.Request, commandChannel chan command) {
	var m members

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	// decode body met json decoder
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Send the ADD command to the channel
	commandChannel <- command{ADD, m.Nodes}

	// write all members as json to the response
	b, me := store.getMembersAsJSON()

	if me != nil {
		log.Printf("Error creating response: " + me.Error())
	} else {
		fmt.Fprintf(w, "%s\n", b)
	}
}

func main() {

	// This is the command channel with a buffer size of 100
	commandChannel := make(chan command, 100)

	go broadcast(commandChannel)
	go processLoop(commandChannel)

	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleGossip(w, r, commandChannel)
		}
	})

	iport := flag.Int("port", 8080, "port of the daemon")
	iseednode := flag.String("seednode", "NONE", "when you specify a seednode, this node will make itself known to main node")

	flag.Parse()
	port := *iport
	seednode := *iseednode

	listenIP := iptools.GetOutboundIP()
	listenAddr := fmt.Sprintf("%s:%d", listenIP, port)
	myHost = "http://" + listenAddr

	memberlist := []string{myHost}
	if seednode != "NONE" {
		memberlist = append(memberlist, seednode)
	}
	commandChannel <- command{ADD, memberlist}

	log.Printf("Started on port %s", myHost)
	log.Fatal(http.ListenAndServe(listenAddr, nil))

}
