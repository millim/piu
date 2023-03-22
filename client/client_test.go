package client

import "testing"

func TestGetPublicIP(t *testing.T) {
	ip := getPublicIP()
	t.Log(ip)
}
