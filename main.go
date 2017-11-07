package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	chDatabase = make(chan interface{}, 10)
}

func main() {
	go dbDistributor()
	http.HandleFunc("/login", oauthCodeGetter)
	http.HandleFunc("/get/addr", hfGetAddr)
	http.HandleFunc("/get/name", hfGetResult)
	http.HandleFunc("/", hfFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hfFile(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if strings.HasPrefix(r.URL.Path, "/index.") {
		f, err := os.Open("./Frontend/index.html")
		if err != nil {
			log.Println(err.Error())
		}
		reader := bufio.NewReader(f)
		byt, _ := ioutil.ReadAll(reader)
		w.Write(byt)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/sign.") {
		f, _ := os.Open("./Frontend/sign.html")
		reader := bufio.NewReader(f)
		byt, _ := ioutil.ReadAll(reader)
		w.Write(byt)
		return
	}
	w.Write([]byte(`404`))
}
