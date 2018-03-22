package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"os"
	"strconv"
	"time"
	"bytes"
	"strings"
	"github.com/trivento/network/iptools"
)
//{"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
type members struct {
	Nodes []string `json:"nodes"`
}

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

func pinger() {
	log.Println("Starting the pinger")
	for ;true;	{
		m, _ := getMembers()
		logKnownHosts()
		for node := range store {
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
				time.Sleep(5 * time.Second)

			}
		}
		time.Sleep(5 * time.Second)
	}
}

func getMembers()([]byte, error) {
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
	go pinger()
	http.HandleFunc("/members", handler)
	port := 8080
	listenIp := iptools.GetOutboundIP()
	if len(os.Args) >= 2 {
		if os.Args[1] != "seed" {
			addMember(os.Args[1])
		}
	}
	if len(os.Args) == 3 {
		p, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println("Invalid port argument %s, falling back to 8080.", os.Args[3])
		} else {
			port = p
		}
		fmt.Println(os.Args[1])
	}
	listenAddr := fmt.Sprintf("%s:%d", listenIp, port)
	myHost = "http://" + listenAddr
	addMember(myHost)
	log.Printf("Started on port %s", myHost)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
