package main

import (
	"net/http"
)

func hfSign(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("签到成功了～"))
}
