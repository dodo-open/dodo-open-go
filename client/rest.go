package client

import (
	"context"
	"dodo-open-go/errs"
	"dodo-open-go/network"
	"dodo-open-go/version"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

// setupResty setup resty client
func (c *client) setupResty() {
	c.r = resty.New().
		SetTransport(createTransport(nil, 500)).
		SetDebug(c.conf.IsDebug).
		SetTimeout(c.conf.Timeout).
		SetAuthToken(c.conf.Token).
		SetAuthScheme("Bot").
		SetHeader("User-Agent", version.Version()).
		OnAfterResponse(
			func(r *resty.Client, resp *resty.Response) error {
				if !network.IsSuccessResponse(resp.StatusCode()) {
					return errs.New(resp.StatusCode(), string(resp.Body()))
				}
				return nil
			},
		)
}

// request you should create a request object before doing each HTTP request
func (c *client) request(ctx context.Context) *resty.Request {
	return c.r.R().SetContext(ctx)
}

// createTransport customize transport
func createTransport(addr net.Addr, idleConnections int) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	if addr != nil {
		dialer.LocalAddr = addr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          idleConnections,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   idleConnections,
		MaxConnsPerHost:       idleConnections,
	}
}
