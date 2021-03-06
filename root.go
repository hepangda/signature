package main

import (
	"encoding/json"
	"net/http"
	"strings"

	xylinux "./XiyouLinuxAPI"
)

func isUsingLinux(ua string) bool {
	return strings.Contains(ua, "Linux") && !strings.Contains(ua, "Windows") && !strings.Contains(ua, "Android")
}

func hfFetchList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	chDatabase <- queryRank{}
	res := <-chDatabase
	byt, _ := json.Marshal(res)
	w.Write(byt)
}

func hfGetResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	state := r.URL.Query().Get("state")
	if state == "" {
		w.Write([]byte(`{"ok":false,"name":"#c"}`))
		return
	}
	query := queryState{
		State: state,
	}

	chDatabase <- query
	res := <-chDatabase
	byt, _ := json.Marshal(res)
	w.Write(byt)
}

func hfGetAddr(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Vaild bool   `json:"vaild"`
		Addr  string `json:"address"`
	}

	isVaild := isUsingLinux(r.UserAgent())
	res := response{
		Vaild: isVaild,
		Addr:  "#",
	}
	if isVaild {
		res.Addr = xylinux.GenerateAddress(app)
	}

	//本句仅限于测试
	w.Header().Add("Access-Control-Allow-Origin", "*")
	byt, _ := json.Marshal(res)
	w.Write(byt)
}
