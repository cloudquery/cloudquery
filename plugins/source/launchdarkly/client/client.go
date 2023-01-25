package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	ldapi "github.com/launchdarkly/api-client-go/v11"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source

	LDSpec   Spec
	Services *ldapi.APIClient
	Backend  backend.Backend
}

func New(logger zerolog.Logger, sourceSpec specs.Source, ldSpec Spec, services *ldapi.APIClient, bk backend.Backend) *Client {
	c := &Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		LDSpec:     ldSpec,
		Services:   services,
		Backend:    bk,
	}
	return c
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, o source.Options) (schema.ClientMeta, error) {
	ldSpec := &Spec{}
	if err := s.UnmarshalSpec(ldSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal launchdarkly spec: %w", err)
	}

	services, err := getServiceClient(s.Version, ldSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *ldSpec, services, o.Backend)
	return cl, nil
}

func getServiceClient(version string, spec *Spec) (*ldapi.APIClient, error) {
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	cfg := ldapi.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", spec.AccessToken)
	cfg.UserAgent += ", cloudquery/source-launchdarkly " + version
	cfg.HTTPClient = &http.Client{
		Timeout: time.Duration(spec.Timeout) * time.Second,
	}

	return ldapi.NewAPIClient(cfg), nil
}
