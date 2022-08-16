// Package httpclient provide a valid configuration for a production usage
// This package need to be improve and move inside github.com/gofast-pkg organization domain
package httpclient

import (
	"net"
	"net/http"
	"time"
)

// default httpclient configuration
const (
	timeoutInMs                   = 10000
	dialerTimeoutInSecond         = 30
	dialerKeepAliveInSecond       = 30
	maxIdleConns                  = 100
	idleConnTimeoutInSecond       = 10
	tlsHandshakeTimeoutInSecond   = 10
	expectContinueTimeoutInSecond = 1
	maxIdleConnsPerHost           = 100
)

// New return an instance of http.Client ready for production usage
func New() *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeoutInMs) * time.Millisecond,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   dialerTimeoutInSecond * time.Second,
				KeepAlive: dialerKeepAliveInSecond * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          maxIdleConns,
			IdleConnTimeout:       idleConnTimeoutInSecond * time.Second,
			TLSHandshakeTimeout:   tlsHandshakeTimeoutInSecond * time.Second,
			ExpectContinueTimeout: expectContinueTimeoutInSecond * time.Second,
			ForceAttemptHTTP2:     true,
			MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		},
	}
}
