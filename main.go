package main

import (
	"log"
	"net/http"
)

func init() {
	chDatabase = make(chan interface{})
}

func main() {
	go dbDistributor()

	http.HandleFunc("/", hfRoot)
	http.HandleFunc("/login", oauthCodeGetter)
	// http.HandleFunc("/sign", hfSign)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
