package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	loginURL     = "https://sso.xiyoulinux.org/oauth/authorize"
	tokenURL     = "https://sso.xiyoulinux.org/oauth/access_token"
	clientSecret = `SECRET`
	clientID     = "dev_java"
	redirectURI  = `http://localhost:8080/login`
)

type oauthToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn string `json:"expires_in"`
}

type oauthCode struct {
	Code  string
	State string
}

func (p *oauthCode) FromQuery(query url.Values) {
	p.Code = query.Get("code")
	p.State = query.Get("state")
}

func generateState() string {
	p := make([]byte, 6)
	ret := make([]byte, 8)
	_, _ = rand.Read(p)
	// base32.StdEncoding.Encode(ret, p)
	base64.URLEncoding.Encode(ret, p)
	return string(ret)[:8]
}

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
		fmt.Println("Error")
	} else {
		apiPrintUserinfo(token)
	}
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
		fmt.Println("Token: " + token.Token)
		return token, nil
	}
	return oauthToken{}, fmt.Errorf("Code Error: Cannot get token")
}

func apiPrintUserinfo(token oauthToken) {
	eg := "https://api.xiyoulinux.org/me?access_token=" + token.Token
	res, err := http.Get(eg)

	if err != nil {
		fmt.Println("Error")
		return
	}

	if res.StatusCode == http.StatusOK {
		body := res.Body
		bodyByte, _ := ioutil.ReadAll(body)
		defer body.Close()

		res := make(map[string]interface{})
		json.Unmarshal(bodyByte, &res)
		printAny := func(name string) {
			fmt.Print(name + ": ")
			fmt.Println(res[name])
		}
		printAny("name")
		printAny("id")
		printAny("major")
	}
}
