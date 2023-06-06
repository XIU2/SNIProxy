package main

import (
	"net"

	"golang.org/x/net/proxy"
)

func GetDialer(isSocks5 bool) proxy.Dialer {
	if !isSocks5 {
		return &net.Dialer{}
	}
	proxyDialer, err := proxy.SOCKS5("tcp", cfg.SocksAddr, nil, proxy.Direct)
	if err != nil {
		// FIXME: I am shit
		return &net.Dialer{}
	}
	return proxyDialer
}
