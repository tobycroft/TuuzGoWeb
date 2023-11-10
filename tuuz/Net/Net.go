package Net

import (
	"net"
	"net/http"
	"time"
)

var dialer = &net.Dialer{
	Timeout:   5 * time.Second,
	KeepAlive: 0 * time.Second,
	//DualStack: true,
}
var transport = &http.Transport{
	DialContext:  dialer.DialContext,
	MaxIdleConns: 100,
}

type Net struct {
	Curl
	WsClient
}
