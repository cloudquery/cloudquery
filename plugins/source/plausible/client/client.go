package client

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

const defaultHTTPTimeout = 30

type Client struct {
	logger     zerolog.Logger
	PluginSpec Spec
	Client     *http.Client
	Backend    backend.Backend
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, options source.Options) (schema.ClientMeta, error) {
	cqClient := Client{
		logger: logger,
	}
	var sfSpec Spec
	if err := spec.UnmarshalSpec(&sfSpec); err != nil {
		return nil, err
	}
	sfSpec.SetDefaults()
	if err := sfSpec.Validate(); err != nil {
		return nil, err
	}
	cqClient.Client = &http.Client{
		Timeout: time.Duration(defaultHTTPTimeout) * time.Second,
	}
	cqClient.PluginSpec = sfSpec
	if options.Backend != nil {
		cqClient.Backend = options.Backend
	}

	return &cqClient, nil
}

func (c *Client) ID() string {
	return c.PluginSpec.SiteId
}
