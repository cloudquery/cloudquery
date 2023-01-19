package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
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
}

func New(logger zerolog.Logger, sourceSpec specs.Source, spSpec Spec, services *shopify.Client) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		spSpec:     spSpec,
		Services:   services,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	spSpec := &Spec{}
	if err := s.UnmarshalSpec(spSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal shopify spec: %w", err)
	}

	services, err := getServiceClient(logger, spSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *spSpec, services)
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

	return shopify.New(
		logger,
		&http.Client{
			Timeout: time.Duration(spec.Timeout) * time.Second,
		},
		spec.APIKey,
		spec.APISecret,
		spec.AccessToken,
		spec.ShopURL,
		spec.MaxRetries,
	)
}
