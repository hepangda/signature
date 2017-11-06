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

//OauthToken 代表一个Oauth 2.0的Access_Token的结构体
//Token 表示access_token的字符串
//Type 代表此Token的类型
//ExpiresIn 代表此Token的失效时间
type OauthToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn string `json:"expires_in"`
}

//OauthCode 代表一个Oauth 2.0登录成功返回的Code
//Code 表示Code的字符串
//State 表示服务器返回的State
type OauthCode struct {
	Code  string
	State string
}

//GenerateAddress 用Secret作为参数执行GenerateAddressS函数，详见该函数的注释
func GenerateAddress(secret Secret) string {
	return GenerateAddressS(secret.ClientID, secret.RedirectURI)
}

//GenerateAddressS 给定指定的客户端ID和转向地址，生成一份随机地该APP登录界面的地址
//clientId 请传入App的client_id值
//redirectURI 请传入App的redirect_uri，当认证成功后将返回到这个页面，该值务必与App在认证服务器上的相同
func GenerateAddressS(clientID string, redirectURI string) string {
	u, _ := url.Parse(URLLogin)
	query := u.Query()

	query.Add("response_type", "code")
	query.Add("client_id", clientID)
	query.Add("redirect_uri", redirectURI)
	query.Add("scope", "all")
	query.Add("state", GenerateState())

	return u.String() + "?" + query.Encode()
}

//GenerateState 将返回一个随机生成的8位的字符串
func GenerateState() string {
	origin := make([]byte, 6)
	coded := make([]byte, 8)
	_, _ = rand.Read(origin)
	base64.URLEncoding.Encode(coded, origin)
	return string(coded)[:8]
}

//GetToken 用Secret作为参数执行GetTokenS函数，详情请见该函数的注解
func GetToken(code OauthCode, secret Secret) (OauthToken, error) {
	return GetTokenS(code, secret.ClientID, secret.ClientSecret, secret.RedirectURI)
}

//GetTokenS 使用Code作为凭证获取到Token
//code 表示从认证服务器获取到的Token
//clientId 请传入App的client_id值
//clientSecret 请传入App的client_secret值，该值由认证服务器提供给你
//redirectURI 请传入App的redirect_uri，当认证成功后将返回到这个页面，该值务必与App在认证服务器上的相同
func GetTokenS(code OauthCode, clientID string, clientSecret string, redirectURI string) (OauthToken, error) {
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

//GenerateOauthCode 从网页值链接中生成一个OauthCode结构体类型
func GenerateOauthCode(query url.Values) OauthCode {
	return OauthCode{
		Code:  query.Get("code"),
		State: query.Get("state"),
	}
}
