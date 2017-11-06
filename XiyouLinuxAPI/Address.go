package xylinux

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
