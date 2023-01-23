package client

import (
	"context"
	"errors"
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
	Backend  backend.Backend
}

func New(logger zerolog.Logger, sourceSpec specs.Source, mpSpec Spec, services *mixpanel.Client, bk backend.Backend) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		MPSpec:     mpSpec,
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
	mpSpec := &Spec{}
	if err := s.UnmarshalSpec(mpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mixpanel spec: %w", err)
	}

	services, err := getServiceClient(logger, mpSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *mpSpec, services, o.Backend)
	return &cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec) (*mixpanel.Client, error) {
	if spec.Secret == "" {
		// allow empty username for the deprecated project secret method: https://developer.mixpanel.com/reference/project-secret
		return nil, errors.New("no credentials provided")
	}
	if spec.ProjectID < 1 {
		return nil, errors.New("no project id provided")
	}

	const dateFormat = "2006-01-02"

	if spec.StartDate == "" {
		dt := time.Now().UTC().AddDate(0, 0, -30).Format(dateFormat)
		logger.Info().Str("start_date", dt).Msg("no start date provided, defaulting to 30 days ago")
		spec.StartDate = dt
	}
	if spec.EndDate == "" {
		dt := time.Now().UTC().AddDate(0, 0, -1).Format(dateFormat)
		logger.Info().Str("end_date", dt).Msg("no end date provided, defaulting to yesterday")
		spec.EndDate = dt
	}

	if _, err := time.Parse(dateFormat, spec.StartDate); err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}
	if _, err := time.Parse(dateFormat, spec.EndDate); err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	if spec.Timeout < 1 {
		spec.Timeout = 30
	}
	if spec.MaxRetries < 1 {
		spec.MaxRetries = 5
	}

	rg, err := mixpanel.ParseRegion(spec.Region)
	if err != nil {
		return nil, err
	}
	if rg == mixpanel.RegionNone {
		rg = mixpanel.RegionUS
	}

	return mixpanel.New(
		mixpanel.ClientOptions{
			Logger: logger,
			HC: &http.Client{
				Timeout: time.Duration(spec.Timeout) * time.Second,
			},
			Region:     rg,
			APIUser:    spec.Username,
			APISecret:  spec.Secret,
			ProjectID:  spec.ProjectID,
			MaxRetries: spec.MaxRetries,
		}), nil
}
