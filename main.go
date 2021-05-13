package main

import (
	"fmt"
	"net"
	"time"
	"flag"
)

func main() {

	hostPtr := flag.String("host", "localhost", "a String")
	protocolPtr := flag.String("protocol", "tcp", "a String")

	flag.parse()

	raw_connect(*hostPtr, flag.Args(), *protocolPtr)
}

func raw_connect(host string, ports []string, protocol string) {
	for_, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout(protocol, net.JoinHostPort(host, port), timeout)
		if err, != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
		}
	}
}