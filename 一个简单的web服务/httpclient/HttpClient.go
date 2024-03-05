package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"net"
	"net/http"
	"syscall"
)

var localIP string
var url string
var interfaceName string

func init() {
	// 给 localIP 参数设置默认值为 "127.0.0.1"，并添加描述信息
	flag.StringVar(&localIP, "localIP", "10.75.52.99", "local IP address to use")
	flag.StringVar(&url, "url", "http://www.baidu.com", "url")
	flag.StringVar(&interfaceName, "interfaceName", "en0", "url")
}

func main() {
	addr, err2 := getInterfaceIpv4Addr("en0")
	fmt.Printf(addr, err2)

	flag.Parse() // 解析参数
	resp, err := HTTPGet(url, localIP)
	if err != nil {
		// 如果有错误，打印出来
		fmt.Println("Error:", err)
		return
	}
	// 首先确保主体在处理完毕后关闭
	defer resp.Body.Close()

	// 用 ioutil.ReadAll 读取响应主体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		// 如果读取时有错误，打印并返回
		return
	}

	// 将读取到的主体 bytes 转换为字符串
	bodyString := string(bodyBytes)

	// 打印状态码，响应主体和可能的错误
	fmt.Printf("Status: %s, Body: %s, Error: %v\n", resp.Status, bodyString, err)
}

//
//// localIP是网卡IP
//func HTTPGet(url, localIP string) (*http.Response, error) {
//	req, _ := http.NewRequest("GET", url, nil)
//	client := &http.Client{
//		Transport: &http.Transport{
//			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
//				// localIP 网卡IP，":0" 表示端口自动选择
//				lAddr, err := net.ResolveTCPAddr(netw, localIP+":0")
//				if err != nil {
//					return nil, err
//				}
//
//				rAddr, err := net.ResolveTCPAddr(netw, addr)
//				if err != nil {
//					return nil, err
//				}
//				conn, err := net.DialTCP(netw, lAddr, rAddr)
//				if err != nil {
//					return nil, err
//				}
//				return conn, nil
//			},
//		},
//	}
//	return client.Do(req)
//}

// localIP是网卡IP
func HTTPGet(url, localIP string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				dialer := net.Dialer{Control: func(network, address string, c syscall.RawConn) error {
					fmt.Printf("DialContext addr:%s,dialer address:%s,network:%s,interfaceName:%s \n", addr, address, network, interfaceName)
					if interfaceName == "" || !isTCPSocket(network) && !isUDPSocket(network) {
						fmt.Println("return nil")
						return nil
					}
					var innerErr error
					err := c.Control(func(fd uintptr) {
						host, _, _ := net.SplitHostPort(address)
						if ip := net.ParseIP(host); ip != nil && !ip.IsGlobalUnicast() {
							fmt.Println("return nil address")
							return
						}
						if interfaceName != "" {
							if innerErr = unix.BindToDevice(int(fd), interfaceName); innerErr != nil {
								fmt.Printf("BindToDevice err:%s", innerErr.Error())
								return
							}
						}
					})
					if innerErr != nil {
						err = innerErr
					}
					return err
				}}
				fmt.Println("return success")
				return dialer.DialContext(ctx, netw, addr)
				//
				//// localIP 网卡IP，":0" 表示端口自动选择
				//lAddr, err := net.ResolveTCPAddr(netw, localIP+":0")
				//if err != nil {
				//	return nil, err
				//}
				//
				//rAddr, err := net.ResolveTCPAddr(netw, addr)
				//if err != nil {
				//	return nil, err
				//}
				//conn, err := net.DialTCP(netw, lAddr, rAddr)
				//if err != nil {
				//	return nil, err
				//}
				//return conn, nil
			},
		},
	}
	return client.Do(req)
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
func getInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	// 通过网卡名获取网络接口信息
	interfaceInfo, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}

	// 通过网络接口信息获取 IP 地址
	addrs, err := interfaceInfo.Addrs()
	if err != nil {
		return "", err
	}

	// 打印网卡的每个 IP 地址
	for _, addr := range addrs {
		if ipv4Addr := addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			return ipv4Addr.String(), nil
		}
	}
	return "", errors.New("getInterfaceIpv4Addr not found")
}
