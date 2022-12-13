package client

import (
	"fmt"
	"os"

	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Spec struct {
	// APIKey can be set via spec or via "TAILSCALE_API_KEY" environment variable
	APIKey string `json:"api_key,omitempty"`

	// Tailnet can be set via spec or via "TAILSCALE_TAILNET" environment variable
	Tailnet string `json:"tailnet,omitempty"`

	// EndpointURL is optional parameter to override the API URL for tailscale.Client.
	EndpointURL string `json:"endpoint_url,omitempty"`
}

const (
	EnvTailscaleApiKey  = "TAILSCALE_API_KEY"
	EnvTailscaleTailnet = "TAILSCALE_TAILNET"
)

func (s *Spec) getClient() (*tailscale.Client, error) {
	if len(s.APIKey) == 0 {
		s.APIKey = os.Getenv(EnvTailscaleApiKey)
		if len(s.APIKey) == 0 {
			return nil, fmt.Errorf("missing API key")
		}
	}
	if len(s.Tailnet) == 0 {
		s.Tailnet = os.Getenv(EnvTailscaleTailnet)
		if len(s.Tailnet) == 0 {
			return nil, fmt.Errorf("missing tailnet")
		}
	}

	var options []tailscale.ClientOption
	if len(s.EndpointURL) > 0 {
		options = append(options, tailscale.WithBaseURL(s.EndpointURL))
	}

	return tailscale.NewClient(s.APIKey, s.Tailnet, options...)
}
