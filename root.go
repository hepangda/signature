package main

import (
	"net/http"
	"strings"
	"time"
)

func hfRoot(w http.ResponseWriter, r *http.Request) {
	isVaild := isFirstTime(r.Cookies()) && isUsingLinux(r.UserAgent())
	if !isVaild {
		w.Write([]byte(`You cannot Visit it!`))
		return
	}

	visitOnce := http.Cookie{
		Name:    "Identify-ID",
		Value:   generateState(),
		Domain:  "localhost",
		Path:    "/",
		Expires: time.Now().Add(5 * time.Second),
	}
	http.SetCookie(w, &visitOnce)
	w.Write([]byte(`OK, and checked!`))
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
