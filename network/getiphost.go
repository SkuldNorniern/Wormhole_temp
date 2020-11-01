package network

import (
	"fmt"
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip := networkIp.IP.String()
			fmt.Println("Resolved Host IP: " + ip)
			return ip
		}
	}
	return ""
}
