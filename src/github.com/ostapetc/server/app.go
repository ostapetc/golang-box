package main

import (
	"net/http"
	"fmt"
	"log"
)

const PORT = ":4000";

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}