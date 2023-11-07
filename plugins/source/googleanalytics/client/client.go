package client

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

type Client struct {
	service *analyticsdata.Service
	Spec    Spec
	backend state.Client

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, spec Spec, backend state.Client) (*Client, error) {
	opts := []option.ClientOption{
		option.WithScopes(analyticsdata.AnalyticsReadonlyScope),
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the safe side with telemetry
		option.WithTelemetryDisabled(),
	}

	if spec.OAuth != nil {
		tokenSource, err := spec.OAuth.getTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	svc, err := analyticsdata.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-googleanalytics"

	return &Client{
		logger:  logger,
		Spec:    spec,
		backend: backend,
		service: svc,
	}, nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "googleanalytics:property-id:{" + c.Spec.PropertyID + "}"
}
