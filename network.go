package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/trivento/network/iptools"
	"math/rand"
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

// Data structuur
//{"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
type members struct {
	Nodes []string `json:"nodes"`
}

// Maak een store waarin alle members gezet worden. Tip: Go kent geen Set
var store = make(map[string]bool)

func addMember(member string) {
	store[member] = true
}

var myHost string

func logKnownHosts() {
	log.Printf("Known hosts:\n")
	for node := range store {
		log.Printf("\t- %s\n", node)
	}
}

// De gossip daemon moet oneindig lang draaien en periodiek een gossip doen naar een aantal, of alle members
func gossipDaemon() {
	log.Println("Starting the gossipDaemon")
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

func getMembers() ([]byte, error) {
	var result []string

	for k := range store {
		result = append(result, k)
	}

	return json.Marshal(members{result})
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var m members

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for _, x := range m.Nodes {
		addMember(x)
	}

	var result []string

	for k := range store {
		result = append(result, k)
	}

	b, me := json.Marshal(members{result})
	if me != nil {
		log.Printf("Error creating response: " + me.Error())
	} else {
		fmt.Fprintf(w, "%s\n", b)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlePost(w, r)
	}
}

func main() {
	go gossipDaemon()
	http.HandleFunc("/members", handler)
	port := 8080
	listenIP := iptools.GetOutboundIP()
	if len(os.Args) >= 2 {
		if os.Args[1] != "seed" {
			addMember(os.Args[1])
		}
	}
	if len(os.Args) == 3 {
		p, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("Invalid port argument %s, falling back to 8080.", os.Args[3])
		} else {
			port = p
		}
		fmt.Println(os.Args[1])
	}
	listenAddr := fmt.Sprintf("%s:%d", listenIP, port)
	myHost = "http://" + listenAddr
	addMember(myHost)
	log.Printf("Started on port %s", myHost)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
