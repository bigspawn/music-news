package internal

import (
	"net"
	"time"

	"golang.org/x/net/proxy"
)

func NewDialer() proxy.Dialer {
	return &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
}
