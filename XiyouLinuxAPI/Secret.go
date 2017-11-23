package xylinux

//Secret 代表一个App相关的标识值
type Secret struct {
	ClientSecret string
	ClientID     string
	RedirectURI  string
}
