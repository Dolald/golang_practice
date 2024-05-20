package main

import (
	"log"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello handler")
	w.WriteHeader(http.StatusOK)
}
