package main

import (
	"log"
	"net/http"
	"strconv"

	xylinux "./XiyouLinuxAPI"
)

/* 以下信息应被在Secret.go文件中定义
const (
	clientSecret   = `Client Secret`
	clientID       = `Client ID`
	redirectURI    = `Redirect URI`
	databaseString = "Username:Password@tcp(127.0.0.1:3306)/signature"

)
*/

func oauthCodeGetter(w http.ResponseWriter, r *http.Request) {
	ret := xylinux.GenerateOauthCode(r.URL.Query())

	token, err := xylinux.GetToken(ret, clientID, clientSecret, redirectURI)
	if err != nil {
		log.Println("Error occured:" + err.Error())
	}
	usr, err := xylinux.GetMessage(token)
	if err != nil {
		log.Println("Error occured:" + err.Error())
	}
	rec := actRecord{
		Dat: signature{
			ID:       strconv.Itoa(usr.ID),
			Username: usr.Name,
		},
	}
	chDatabase <- rec
	http.Redirect(w, r, "../sign", http.StatusFound)
}
