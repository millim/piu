package client

import (
	"flag"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Time      string
	ServerUrl string
}

func (c *Client) InitArgs() {
	flag.StringVar(&c.Time, "time", "30m", "duration time, example: 1h, 3m, default: 30m")
	flag.StringVar(&c.ServerUrl, "server-url", "http://localhost:8080", "server url")
}

func (c *Client) Run() {
	duration, err := time.ParseDuration(c.Time)
	if err != nil {
		fmt.Println("timeFormat error, example: 10s, 1h")
	}
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			updateIPs()
		}
	}()
	select {}
}

func getLocalIPs() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To16() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

func updateIPs() {
	ips := getLocalIPs()
	fmt.Println("ips -->", ips)
}
