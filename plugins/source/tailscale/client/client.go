package client

import (
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
	Tailnet string
	APIKey  string
	Clients map[string]*tailscale.Client
}

func (c *Client) WithTailNet(tailnet string) *Client {
	client := *c
	client.Tailnet = tailnet
	return &client
}
