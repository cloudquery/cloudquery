package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	metrics    destination.Metrics
	pluginSpec Spec
	client     *bigquery.Client
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {
	if destSpec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("bigquery destination only supports append mode")
	}
	var err error
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

	// the context here is used for token refresh so this is workaround as suggested
	// https://github.com/googleapis/google-cloud-go/issues/946
	// https://github.com/googleapis/google-cloud-go/commit/2d59af0cb37fb29e5b7980a15088938778f117c7
	c.client, err = c.bqClient(context.Background())
	if err != nil {
		return nil, err
	}

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

func (c *Client) Close(_ context.Context) error {
	return c.client.Close()
}
