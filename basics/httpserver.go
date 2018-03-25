package main

import (
	"fmt"
	"log"
	"net/http"
)

func startHTTPServer() {

	http.HandleFunc("/rest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("received %v\n", r)
		fmt.Fprintf(w, "Running\n")
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
