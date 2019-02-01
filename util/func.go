package util

import (
	"net/http"
	"net"
	"strings"
	"fmt"
	"github.com/jecqiang/mygo"
)

//获取客户端IP
func GetClientIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr

}

// 获取客户端访问的浏览器
func GetClientBrowser(req *http.Request) string {
	browser := "unknown"
	userAgen := req.Header.Get("user-agent")
	if strings.Contains(userAgen, "Chrome") {
		browser = "Chrome"
	} else if strings.Contains(userAgen, "Firefox") {
		browser = "Firefox"
	} else if strings.Contains(userAgen, "MSIE") {
		browser = "IE"
	} else if strings.Contains(userAgen, "OPR") || strings.Contains(userAgen, "Opera") {
		browser = "Opera"
	} else if strings.Contains(userAgen, "Edge") {
		browser = "Edge"
	} else if strings.Contains(userAgen, "Safari") {
		browser = "Safari"
	} else if strings.Contains(userAgen, "360SE") {
		browser = "360SE"
	}
	return browser

}

func MailTest() {
	to := []string{"zhouzeqiang@banggood.com"}
	cc := []string{"jecqiang@163.com"}
	msg := &mygo.MailMessage{
		To:      to,
		Cc:      cc,
		Subject: "test",
		Body:    "body",
	}
	err := mygo.G_mail.Send(msg)
	fmt.Println(err)
}
