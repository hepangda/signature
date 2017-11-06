package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/url"
)

type signature struct {
	ID       string `json:"id"`
	Username string `json:"name"`
}

type action interface {
	Do(conn *sql.DB) error
}

type actRecord struct {
	Dat signature
}

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
	base64.URLEncoding.Encode(ret, p)
	return string(ret)[:8]
}
