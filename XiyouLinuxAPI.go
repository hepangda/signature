package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// XlgUser 存储一个XiyouLinux用户信息
// 目前尚不能获得的数据项有：
// online_at updated_at native created_at group
type XlgUser struct {
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

func apiGetMessage(token oauthToken) (XlgUser, error) {
	eg := URLGetMe + "?access_token=" + token.Token
	res, err := http.Get(eg)

	if err != nil {
		return XlgUser{}, fmt.Errorf("Error")
	}

	if res.StatusCode == http.StatusOK {
		body := res.Body
		bodyByte, _ := ioutil.ReadAll(body)
		defer body.Close()

		var res XlgUser
		json.Unmarshal(bodyByte, &res)

		return res, nil
	}
	return XlgUser{}, fmt.Errorf("Bad Request")
}
