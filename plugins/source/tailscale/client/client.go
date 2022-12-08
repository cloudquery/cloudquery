package client

import (
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
	Tailscale *tailscale.Client
}
