package main

import (
	"fmt"
	"net/http"
)

func portOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK - Port One")
}

func portTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK - Port Two")
}

func main() {
	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/healthz", portOne)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/healthz", portTwo)

	go func() {
		http.ListenAndServe(":8081", serverMuxA)
	}()

	http.ListenAndServe(":8082", serverMuxB)
}
