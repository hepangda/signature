package main

import (
	"net/http"
	"strings"
	"time"

	xylinux "./XiyouLinuxAPI"
)

func hfRoot(w http.ResponseWriter, r *http.Request) {
	if !isUsingLinux(r.UserAgent()) {
		w.Write([]byte(`请检查你是否正在使用Linux系统登录。`))
		return
	}

	if !isFirstTime(r.Cookies()) {
		w.Write([]byte(`您已经签到过了，你可以点击此处查看您的签到信息（暂未支持）`))
		return
	}

	visitOnce := http.Cookie{
		Name:    "Identify-ID",
		Value:   xylinux.GenerateState(),
		Domain:  "localhost",
		Path:    "/",
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, &visitOnce)
	http.Redirect(w, r, xylinux.GenerateAddress(clientID, redirectURI), http.StatusFound)
}

func isFirstTime(cookies []*http.Cookie) bool {
	for _, i := range cookies {
		if i.Name == "Identify-ID" {
			return false
		}
	}
	return true
}

func isUsingLinux(ua string) bool {
	return strings.Contains(ua, "Linux") && !strings.Contains(ua, "Windows") && !strings.Contains(ua, "Android")
}
