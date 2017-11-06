package xylinux

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//GetMessage 通过Token获取到该用户的相关信息
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
