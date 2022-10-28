package main

import (
	"fmt"
	"net/http"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", FirstHandler)
	http.ListenAndServe(":8080", nil)
}