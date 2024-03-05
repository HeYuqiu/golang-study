package main

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"net"
	"net/http"
	"net/url"
	"syscall"
)

var interfaceName string = "ens160" // ens33、ens160
const (
	network string = "tcp"
	address string = "192.168.3.58:8080"
)

func main() {
	// 重要核心代码
	d := &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return setSocketOptions(network, address, c, interfaceName)
		},
	}
	ctx := context.Background()
	// 拨号获取一个连接
	conn, err := d.DialContext(ctx, network, address)
	if err != nil {
		panic(err)
	}
	// 1. create new request
	reqURL, _ := url.Parse("http://web.test.com/test/index.html")
	hdr := http.Header{}
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: hdr,
	}
	err = req.Write(conn)
	if err != nil {
		panic(err)
	}
	// 2. Get a response
	r := bufio.NewReader(conn)
	resp, err := http.ReadResponse(r, req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.LocalAddr())
	for name, vv := range resp.Header {
		fmt.Print(name, ":")
		for _, v := range vv {
			fmt.Print(v)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println(status)
	fmt.Println(string(body))
}
func isTCPSocket(network string) bool {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return true
	default:
		return false
	}
}
func isUDPSocket(network string) bool {
	switch network {
	case "udp", "udp4", "udp6":
		return true
	default:
		return false
	}
}
func setSocketOptions(network, address string, c syscall.RawConn, interfaceName string) (err error) {
	if interfaceName == "" || !isTCPSocket(network) && !isUDPSocket(network) {
		return
	}
	var innerErr error
	err = c.Control(func(fd uintptr) {
		host, _, _ := net.SplitHostPort(address)
		if ip := net.ParseIP(host); ip != nil && !ip.IsGlobalUnicast() {
			return
		}
		if interfaceName != "" {
			if innerErr = unix.BindToDevice(int(fd), interfaceName); innerErr != nil {
				return
			}
		}
	})
	if innerErr != nil {
		err = innerErr
	}
	return
}
