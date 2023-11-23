package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping",  func(http.ResponseWriter, *http.Request) {
		log.Println("ping")
	})
	
	http.ListenAndServe(":8080", nil)
}