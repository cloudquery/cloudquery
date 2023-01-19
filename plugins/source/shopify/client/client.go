package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	spSpec     Spec

	Services *shopify.Client
	Backend  backend.Backend
}

func New(logger zerolog.Logger, sourceSpec specs.Source, spSpec Spec, services *shopify.Client, bk backend.Backend) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		spSpec:     spSpec,
		Services:   services,
		Backend:    bk,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, o source.Options) (schema.ClientMeta, error) {
	spSpec := &Spec{}
	if err := s.UnmarshalSpec(spSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal shopify spec: %w", err)
	}

	services, err := getServiceClient(logger, spSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *spSpec, services, o.Backend)
	return &cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec) (*shopify.Client, error) {
	if spec.AccessToken == "" && (spec.APIKey == "" || spec.APISecret == "") {
		return nil, errors.New("no credentials provided")
	}
	if spec.ShopURL == "" {
		return nil, errors.New("no shop url provided")
	}
	if !strings.HasSuffix(spec.ShopURL, ".myshopify.com") {
		return nil, errors.New("shop url should end with .myshopify.com, as in https://shop_name.myshopify.com")
	}

	if spec.Timeout < 1 {
		spec.Timeout = 10
	}
	if spec.MaxRetries < 1 {
		spec.MaxRetries = 30
	}
	if spec.PageSize < 1 {
		spec.PageSize = 50
	}

	return shopify.New(shopify.ClientOptions{
		Log: logger,
		HC: &http.Client{
			Timeout: time.Duration(spec.Timeout) * time.Second,
		},
		ApiKey:      spec.APIKey,
		ApiSecret:   spec.APISecret,
		AccessToken: spec.AccessToken,
		ShopURL:     spec.ShopURL,
		MaxRetries:  spec.MaxRetries,
		PageSize:    int(spec.PageSize),
	})
}
