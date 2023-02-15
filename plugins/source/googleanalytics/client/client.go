package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

type Client struct {
	service *analyticsdata.Service
	backend backend.Backend

	reports []*Report

	PropertyID string
	StartDate  string

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "googleanalytics:property-id:{" + c.PropertyID + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	spec := new(Spec)
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	spec.setDefaults()
	if err := spec.validate(); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		option.WithScopes(analyticsdata.AnalyticsReadonlyScope),
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemetry
		option.WithTelemetryDisabled(),
	}

	if spec.OAuth != nil {
		tokenSource, err := spec.OAuth.getTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	svc, err := analyticsdata.NewService(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-googleanalytics/" + srcSpec.Version

	c := &Client{
		service:    svc,
		backend:    options.Backend,
		StartDate:  spec.StartDate,
		PropertyID: spec.PropertyID,
		reports:    spec.Reports,
		logger: logger.With().
			Str("plugin", "googleanalytics").
			Str("property_id", spec.PropertyID).
			Logger(),
	}

	return c, nil
}
