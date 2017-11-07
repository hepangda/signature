package main

import (
	"log"
	"net/http"
)

func init() {
	chDatabase = make(chan interface{}, 10)
}

func main() {
	go dbDistributor()
	http.HandleFunc("/login", oauthCodeGetter)
	http.HandleFunc("/get/addr", hfGetAddr)
	http.HandleFunc("/get/name", hfGetResult)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
