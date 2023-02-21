package ipx

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP 获取客户端ip
func ClientIP(req *http.Request) (ip string) {
	for _, ip = range strings.Split(req.Header.Get("X-Forwarded-For"), ",") {
		if !IsLocalIPWithStr(strings.TrimSpace(ip)) {
			return
		}
	}
	if !IsLocalIPWithStr(strings.TrimSpace(req.Header.Get("X-Real-Ip"))) {
		return
	}
	if !IsLocalIPWithStr(RemoteIP(req)) {
		return
	}
	return
}

func RemoteIP(req *http.Request) (ip string) {
	ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	return
}
