package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
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

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	pluginSpec := new(Spec)
	if err := spec.UnmarshalSpec(pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	if err := pluginSpec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	// using the new oauth mechanism
	if pluginSpec.APIKey == "" {
		oatuhURL := "https://api.tailscale.com/api/v2/oauth/token"
		resp, err := http.DefaultClient.PostForm("https://api.tailscale.com/api/v2/oauth/token", url.Values{
			"client_id":     {pluginSpec.ClientID},
			"client_secret": {pluginSpec.ClientSecret},
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

		pluginSpec.APIKey = res.AccessToken
	}

	var options []tailscale.ClientOption
	if len(pluginSpec.EndpointURL) > 0 {
		options = append(options, tailscale.WithBaseURL(pluginSpec.EndpointURL))
	}

	tailscaleClient, err := tailscale.NewClient(pluginSpec.APIKey, pluginSpec.Tailnet, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create tailscale client: %w", err)
	}

	return &Client{
		TailscaleClient: tailscaleClient,
		pluginSpec:      pluginSpec,
		Logger:          logger,
	}, nil
}
