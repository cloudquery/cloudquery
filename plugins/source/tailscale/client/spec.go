package client

import (
	"fmt"

	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Spec struct {
	APIKey      string `json:"api_key,omitempty"`
	Tailnet     string `json:"tailnet,omitempty"`
	EndpointURL string `json:"endpoint_url,omitempty"`
}

func (s *Spec) getClient() (*tailscale.Client, error) {
	if len(s.APIKey) == 0 {
		return nil, fmt.Errorf("missing API key")
	}
	if len(s.Tailnet) == 0 {
		return nil, fmt.Errorf("missing tailnet")
	}

	var options []tailscale.ClientOption
	if len(s.EndpointURL) > 0 {
		options = append(options, tailscale.WithBaseURL(s.EndpointURL))
	}

	return tailscale.NewClient(s.APIKey, s.Tailnet, options...)
}
