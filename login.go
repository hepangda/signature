package main

import (
	"log"
	"net/http"
	"strconv"

	xylinux "./XiyouLinuxAPI"
)

/* 以下信息应被在Secret.go文件中定义
var app = xylinux.Secret{
	ClientSecret: `client_secret`,
	ClientID:     "client_id",
	RedirectURI:  `redirect_uri`,
}

const databaseString = "username:password@tcp(127.0.0.1:3306)/signature"
*/

func oauthCodeGetter(w http.ResponseWriter, r *http.Request) {
	ret := xylinux.GenerateOauthCode(r.URL.Query())

	token, err := xylinux.GetToken(ret, app)
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
			State:    ret.State,
		},
	}
	chDatabase <- rec
	http.Redirect(w, r, `file:///home/pangda/GoProjects/signature/sign.html?state=`+ret.State, http.StatusFound)
}
