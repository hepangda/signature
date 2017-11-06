package main

import (
	"log"
	"net/http"

	xylinux "./XiyouLinuxAPI"
)

func init() {
	chDatabase = make(chan interface{})
}

func main() {
	go dbDistributor()
	xylinux.Test()
	http.HandleFunc("/", hfRoot)
	http.HandleFunc("/login", oauthCodeGetter)
	http.HandleFunc("/sign", hfSign)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
