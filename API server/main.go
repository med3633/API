package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page.\n")
	fmt.Println("Endpoint hit: homepage")
}

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)
}
