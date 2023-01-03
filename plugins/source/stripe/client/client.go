package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	stSpec     Spec

	Services *client.API
}

func New(logger zerolog.Logger, sourceSpec specs.Source, stSpec Spec, services *client.API) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		stSpec:     stSpec,
		Services:   services,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	stSpec := &Spec{}
	if err := s.UnmarshalSpec(stSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stripe spec: %w", err)
	}

	services, err := getServiceClient(logger, stSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *stSpec, services)
	return &cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec) (*client.API, error) {
	if spec.APIKey == "" {
		return nil, errors.New("no api_key provided")
	}

	sCfg := &stripe.BackendConfig{
		LeveledLogger: &LeveledLogger{
			Logger: logger.With().Str("source", "stripe-client").Logger(),
		},
	}
	if spec.MaxRetries > 0 {
		sCfg.MaxNetworkRetries = stripe.Int64(spec.MaxRetries)
	}

	c := &client.API{}
	c.Init(spec.APIKey, &stripe.Backends{
		API:     stripe.GetBackendWithConfig(stripe.APIBackend, sCfg),
		Connect: stripe.GetBackendWithConfig(stripe.ConnectBackend, sCfg),
		Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, sCfg),
	})
	return c, nil
}
