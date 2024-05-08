package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", middleWare(handler))

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello handler")
	w.WriteHeader(http.StatusOK)
}

func middleWare(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello middle ware handler")
		nextHandler(w, r)
	}
}
