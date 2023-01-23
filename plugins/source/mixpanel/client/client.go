package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source

	MPSpec   Spec
	Services *mixpanel.Client
	Backend  *BackendWrapper
}

func New(logger zerolog.Logger, sourceSpec specs.Source, mpSpec Spec, services *mixpanel.Client, bk backend.Backend) *Client {
	c := &Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		MPSpec:     mpSpec,
		Services:   services,
	}
	if bk != nil {
		c.Backend = NewBackendWrapper(bk)
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
	mpSpec := &Spec{}
	if err := s.UnmarshalSpec(mpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mixpanel spec: %w", err)
	}

	services, err := getServiceClient(logger, mpSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *mpSpec, services, o.Backend)
	return cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec) (*mixpanel.Client, error) {
	spec.SetDefaults(logger)
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	return mixpanel.New(
		mixpanel.ClientOptions{
			Logger: logger,
			HC: &http.Client{
				Timeout: time.Duration(spec.Timeout) * time.Second,
			},
			Region:     spec.region,
			APIUser:    spec.Username,
			APISecret:  spec.Secret,
			ProjectID:  spec.ProjectID,
			MaxRetries: spec.MaxRetries,
		}), nil
}
