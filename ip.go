package ipx

import "net"

// IsIP 是否为ip
func IsIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}
