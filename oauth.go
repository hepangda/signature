package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

/* 以下信息应被在Secret.go文件中定义
const (
	loginURL     = `URL to Login`
	tokenURL     = `URL to Get Token`
	clientSecret = `Client Secret`
	clientID     = `Client ID`
	redirectURI  = `Redirect URI`
)
*/

func oauthGenerateAddress() string {
	u, _ := url.Parse(loginURL)
	uq := u.Query()
	uq.Add("response_type", "code")
	uq.Add("client_id", clientID)
	uq.Add("redirect_uri", redirectURI)
	uq.Add("scope", "all")
	uq.Add("state", generateState())

	return u.String() + "?" + uq.Encode()
}

func oauthCodeGetter(w http.ResponseWriter, r *http.Request) {
	var ret oauthCode
	ret.FromQuery(r.URL.Query())
	token, err := oauthGetToken(ret)
	if err != nil {
		log.Println("Error occured:" + err.Error())
	}
	usr, err := apiGetMessage(token)
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

func generateParaFromMap(pm map[string]string) []byte {
	var ret string
	for key, value := range pm {
		ret += key + "=" + value + "&"
	}
	return []byte(ret[:len(ret)-1])
}

func oauthGetToken(code oauthCode) (oauthToken, error) {
	paraMap := make(map[string]string)
	paraMap["code"] = code.Code
	paraMap["grant_type"] = "authorization_code"
	paraMap["client_id"] = clientID
	paraMap["client_secret"] = clientSecret
	paraMap["redirect_uri"] = redirectURI

	postBody := bytes.NewBuffer(generateParaFromMap(paraMap))
	res, err := http.Post(tokenURL, "application/x-www-form-urlencoded", postBody)

	if err != nil {
		log.Fatalln("HTTP Fatal: Cannot post.")
	}

	if res.StatusCode == http.StatusCreated {
		body := res.Body
		bodyByte, _ := ioutil.ReadAll(body)
		defer body.Close()

		var token oauthToken
		json.Unmarshal(bodyByte, &token)
		return token, nil
	}
	return oauthToken{}, fmt.Errorf("Code Error: Cannot get token")
}
