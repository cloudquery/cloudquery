package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

type Client struct {
	plugins.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	metrics    plugins.DestinationMetrics
	pluginSpec Spec
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (plugins.DestinationClient, error) {
	if destSpec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("bigquery destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "bq-dest").Logger(),
		spec:   destSpec,
	}
	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BigQuery spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec

	// create a client to test that we can do it, but new clients will also be instantiated
	// for queries so that we can use a new context there.
	client, err := c.bqClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create new BigQuery client: %w", err)
	}
	defer client.Close()

	return c, nil
}

func (c *Client) bqClient(ctx context.Context) (*bigquery.Client, error) {
	opts := []option.ClientOption{option.WithRequestReason("CloudQuery BigQuery destination")}
	if len(c.pluginSpec.ServiceAccountKeyJSON) != 0 {
		opts = append(opts, option.WithCredentialsJSON([]byte(c.pluginSpec.ServiceAccountKeyJSON)))
	}
	client, err := bigquery.NewClient(ctx, c.pluginSpec.ProjectID, opts...)
	if err != nil {
		return nil, err
	}
	if c.pluginSpec.DatasetLocation != "" {
		client.Location = c.pluginSpec.DatasetLocation
	}
	return client, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}
