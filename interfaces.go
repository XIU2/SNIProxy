package main

import (
	"fmt"
	"net"

	"golang.org/x/net/proxy"
)

func GetDialer(isSocks5 bool) proxy.Dialer {
	if !isSocks5 { // 如果配置文件中未开启 SOCKS5 代理，则直接返回直连原接口
		return &net.Dialer{}
	}

	// 当配置了代理账号或密码，那么就需要创建认证对象
	/*var auth *proxy.Auth
	if cfg.SocksUser != "" || cfg.SocksPassword != "" {
		auth = &proxy.Auth{
			User:     cfg.SocksUser,
			Password: cfg.SocksPassword,
		}
	}
	proxyDialer, err := proxy.SOCKS5("tcp", cfg.SocksAddr, auth, proxy.Direct)*/

	proxyDialer, err := proxy.SOCKS5("tcp", cfg.SocksAddr, nil, proxy.Direct)
	if err != nil {
		serviceLogger(fmt.Sprintf("连接 SOCKS5 代理时出错, 已回退为直连, %v", err), 31, false)
		return &net.Dialer{}
	}
	return proxyDialer
}
