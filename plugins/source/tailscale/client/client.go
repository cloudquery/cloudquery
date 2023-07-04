package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
	plugin.UnimplementedDestination
	TailscaleClient *tailscale.Client
	pluginSpec      *Spec
	Logger          zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) ID() string {
	return c.pluginSpec.Tailnet
}

type oauthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiredIn   int    `json:"expired_in"`
}

func Configure(ctx context.Context, logger zerolog.Logger, spec *Spec) (schema.ClientMeta, error) {
	// using the new oauth mechanism
	if spec.APIKey == "" {
		oatuhURL := "https://api.tailscale.com/api/v2/oauth/token"
		resp, err := http.DefaultClient.PostForm("https://api.tailscale.com/api/v2/oauth/token", url.Values{
			"client_id":     {spec.ClientID},
			"client_secret": {spec.ClientSecret},
		})
		if err != nil {
			return nil, fmt.Errorf("error getting keys from %s: %w", oatuhURL, err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body from %s: %w", oatuhURL, err)
		}

		res := oauthResponse{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, fmt.Errorf("error unmarshalling response body from %s: %w", oatuhURL, err)
		}

		spec.APIKey = res.AccessToken
	}

	var options []tailscale.ClientOption
	if len(spec.EndpointURL) > 0 {
		options = append(options, tailscale.WithBaseURL(spec.EndpointURL))
	}

	tailscaleClient, err := tailscale.NewClient(spec.APIKey, spec.Tailnet, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create tailscale client: %w", err)
	}

	return &Client{
		TailscaleClient: tailscaleClient,
		pluginSpec:      spec,
		Logger:          logger,
	}, nil
}
