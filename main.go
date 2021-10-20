package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8099", nil)
	log.Fatal(fmt.Printf("error while serving http: %v", err))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
