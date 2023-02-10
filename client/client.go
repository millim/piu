package client

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
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
	fmt.Println(c.ServerUrl)
	updateIPs(c.ServerUrl)
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			updateIPs(c.ServerUrl)
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

func updateIPs(serverURL string) {
	ips := getLocalIPs()
	body := map[string]string{
		"key":   "ip",
		"value": strings.Join(ips, ","),
	}
	data, _ := json.Marshal(body)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/values", serverURL), bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println("request err:", err)
		return
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client request err:", err)
		return
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("update ip success")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
