package ipx

import (
	"net"
)

// IsLocalIP 是否本地ip
// A类：10.0.0.0 - 10.255.255.255
// B类：172.16.0.0 - 172.31.255.255
// C类：192.168.0.0 - 192.168.255.255
func IsLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip[1] <= 31) ||
		(ip4[0] == 192 && ip4[1] == 168)
}

func IsLocalIPWithStr(ip string) bool {
	return IsLocalIP(net.ParseIP(ip))
}
