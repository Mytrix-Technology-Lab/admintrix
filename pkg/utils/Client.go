package utils

import "github.com/gogf/gf/net/ghttp"

func GetClientIp(req *ghttp.Request) string {
	return ""
}

func GetIpCity(ipclient string) string {
	return ""
}

type UserInfo struct {
	Realname string
	Username string
}

func UInfo(req *ghttp.Request) UserInfo {
	return UserInfo{}
}

func Uid(req *ghttp.Request) int {
	return 0
}
