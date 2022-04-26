package dodo_open_go

import "github.com/dodo-open/dodo-open-go/client"

// NewInstance create a new DoDoBot instance
func NewInstance(clientId, token string, options ...client.OptionHandler) (client.Client, error) {
	return client.New(clientId, token, options...)
}
