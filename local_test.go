package ipx

import (
	"net"
	"testing"
)

func TestIsLocalIP(t *testing.T) {
	list := []struct {
		Ip string
		Is bool
	}{
		{Ip: "10.0.0.0", Is: true},
		{Ip: "10.254.255.255", Is: true},
		{Ip: "10.255.255.255", Is: true},
		{Ip: "172.31.255.255", Is: true},
		{Ip: "172.16.0.0", Is: true},
		{Ip: "172.18.0.0", Is: true},
		{Ip: "172.13.0.0", Is: false},
		{Ip: "192.168.0.0", Is: true},
		{Ip: "192.168.255.0", Is: true},
		{Ip: "192.165.255.0", Is: false},
	}

	for _, s := range list {
		if IsLocalIP(net.ParseIP(s.Ip)) != s.Is {
			t.Error(s.Ip + " no no no")
		}
	}
}
