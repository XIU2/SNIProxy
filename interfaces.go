package main

import (
	"net"

	"golang.org/x/net/proxy"
)

func GetDialer(isSocks5 bool) proxy.Dialer {
	if !isSocks5 { // 如果配置文件中未开启 SOCKS5 代理，则直接返回直连原接口
		return &net.Dialer{}
	}

	// 当配置了代理账号或密码，那么就需要创建认证对象（如果都没有配置，那么 auth 的值就会是 nil，下面 proxy.SOCKS5 也就不会启用身份认证了）
	var auth *proxy.Auth
	if cfg.SocksUsername != "" || cfg.SocksPassword != "" {
		auth = &proxy.Auth{
			User:     cfg.SocksUsername,
			Password: cfg.SocksPassword,
		}
	}
	proxyDialer, err := proxy.SOCKS5("tcp", cfg.SocksAddr, auth, proxy.Direct)
	if err != nil { // 如果报错就退回直连（暂时没找到会在此处引起报错的情况）
		return &net.Dialer{}
	}
	return proxyDialer
}
