package net

import (
	"net"
	"strconv"
	"strings"
	"time"
)

type InterfaceIPAddress struct {
	Iface string
	IP    string
}

// WaitForPort wait for successful network connection
func WaitForPort(proto string, ip string, port int, timeout time.Duration) error {
	for {
		con, err := net.DialTimeout(proto, ip+":"+strconv.Itoa(port), timeout)
		if err == nil {
			con.Close()
			break
		}
	}

	return nil
}

// RandomPort return a random not used TCP port
func RandomPort() string {
	l, _ := net.Listen("tcp", "127.0.0.1:0")
	defer l.Close()
	port := l.Addr()
	return strings.Split(port.String(), ":")[1]
}

// GetNetworkInterfaces return the list of
// network interfaces and IP address
func GetNetworkInterfaces() []InterfaceIPAddress {
	result := []InterfaceIPAddress{}

	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				result = append(result, InterfaceIPAddress{inter.Name, addr.String()})
			}
		}
	}

	return result
}

func ParseIP(s string) net.IP {
	return net.ParseIP(s)
}
