package main

import (
	"context"
	"log"
	"net"
	"net/http"
)

func main() {
	transport := &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		interfaces, err := net.Interfaces()
		if err != nil {
			log.Fatal(err)
		}
		var chosenInterface net.Interface
		for _, iface := range interfaces {
			if iface.Name == "eth0" {
				chosenInterface = iface
				break
			}
		}
		dialer := net.Dialer{LocalAddr: &net.IPAddr{IP: chosenInterface.IP}}
		return dialer.DialContext(ctx, network, addr)
	}}
	client := &http.Client{Transport: transport}
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// xxxx

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	var chosenInterface net.Interface
	for _, iface := range interfaces {
		if iface.Name == "eth0" {
			chosenInterface = iface
			break
		}
	}
	dialer := net.Dialer{LocalAddr: &net.IPAddr{IP: chosenInterface.IP}}
	conn, err := dialer.DialContext(context.Background(), "tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
