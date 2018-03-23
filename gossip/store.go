package main

import (
	"encoding/json"
	"sync"
)

// Maak een store waarin alle members gezet worden. Tip: Go kent geen Set
var (
	store     = make(map[string]bool)
	storelock sync.Mutex
)

// Data structuur
//{"nodes":["http://10.248.30.150:8082","http://10.248.30.150:8081"]}
type members struct {
	Nodes []string `json:"nodes"`
}

func addMembers(memberlist []string) {
	storelock.Lock()
	for _, m := range memberlist {
		store[m] = true
	}
	storelock.Unlock()

}

func getMembers() ([]byte, error) {
	var result []string
	storelock.Lock()

	for k := range store {
		result = append(result, k)
	}
	storelock.Unlock()

	return json.Marshal(members{result})
}
