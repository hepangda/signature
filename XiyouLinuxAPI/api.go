package xylinux

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
	//URLLogin 登录获取Code的接口地址
	URLLogin = "https://sso.xiyoulinux.org/oauth/authorize"
	//URLLogout 登出接口地址，直接对此地址发Get请求即可
	URLLogout = "https://sso.xiyoulinux.org/logout"
	//URLGetToken 根据Code获取Token的接口地址
	URLGetToken = "https://sso.xiyoulinux.org/oauth/access_token"
	//URLGetMe 获取用户信息的API接口地址
	URLGetMe = "https://api.xiyoulinux.org/me"
)

func Test() {
	fmt.Println("XiyouLinuxAPI golang version.")
}

type OauthToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn string `json:"expires_in"`
}

type OauthCode struct {
	Code  string
	State string
}

// User 存储一个XiyouLinux用户信息
// 目前尚不能获得的数据项有：
// online_at updated_at native created_at group
type User struct {
	AvatarURL string `json:"avatar_url"`
	Sex       string `json:"sex"`
	Phone     string `json:"phone"`
	BlogURL   string `json:"blog_url"`
	Job       string `json:"job"`
	Workplace string `json:"workplace"`
	GithubURL string `json:"github_url"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Grade     string `json:"grade"`
	Wechat    string `json:"wechat"`
	QQ        string `json:"qq"`
	Major     string `json:"major"`
}

type HttpParam map[string]string

func GenerateParamStream(param HttpParam) []byte {
	var ret string
	for k, v := range param {
		ret += k + "=" + v + "&"
	}
	return []byte(ret[:len(ret)-1])
}

func GenerateAddress(clientID string, redirectURI string) string {
	u, _ := url.Parse(URLLogin)
	query := u.Query()

	query.Add("response_type", "code")
	query.Add("client_id", clientID)
	query.Add("redirect_uri", redirectURI)
	query.Add("scope", "all")
	query.Add("state", GenerateState())

	return u.String() + "?" + query.Encode()
}

func GenerateState() string {
	origin := make([]byte, 6)
	coded := make([]byte, 8)
	_, _ = rand.Read(origin)
	base64.URLEncoding.Encode(coded, origin)
	return string(coded)[:8]
}

func GenerateOauthCode(query url.Values) OauthCode {
	return OauthCode{
		Code:  query.Get("code"),
		State: query.Get("state"),
	}
}

func GetMessage(token OauthToken) (User, error) {
	eg := URLGetMe + "?access_token=" + token.Token
	res, err := http.Get(eg)

	if err != nil {
		return User{}, fmt.Errorf("Error")
	}

	if res.StatusCode == http.StatusOK {
		body := res.Body
		bodyByte, _ := ioutil.ReadAll(body)
		defer body.Close()

		var res User
		json.Unmarshal(bodyByte, &res)

		return res, nil
	}
	return User{}, fmt.Errorf("Bad Request")
}

func GetToken(code OauthCode, clientID string, clientSecret string, redirectURI string) (OauthToken, error) {
	paraMap := make(map[string]string)
	paraMap["code"] = code.Code
	paraMap["grant_type"] = "authorization_code"
	paraMap["client_id"] = clientID
	paraMap["client_secret"] = clientSecret
	paraMap["redirect_uri"] = redirectURI

	postBody := bytes.NewBuffer(GenerateParamStream(paraMap))
	res, err := http.Post(URLGetToken, "application/x-www-form-urlencoded", postBody)

	if err != nil {
		log.Fatalln("HTTP Fatal: Cannot post.")
	}

	if res.StatusCode == http.StatusCreated {
		body := res.Body
		bodyByte, _ := ioutil.ReadAll(body)
		defer body.Close()

		var token OauthToken
		json.Unmarshal(bodyByte, &token)
		return token, nil
	}
	return OauthToken{}, fmt.Errorf("Code Error: Cannot get token")
}
