package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//APIURLGetMe 获取用户信息的API接口地址
	APIURLGetMe = "https://api.xiyoulinux.org/me"
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
	eg := APIURLGetMe + "?access_token=" + token.Token
	res, err := http.Get(eg)

	if err != nil {
		return XlgUser{}, fmt.Errorf("Error")
	}

	log.Printf("Status: %d\n", res.StatusCode)
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
