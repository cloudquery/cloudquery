package client

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   Spec
	client *bigquery.Client
	writer *writers.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "bq-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BigQuery spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = writers.NewBatchWriter(c, writers.WithLogger(logger), writers.WithBatchSize(c.spec.BatchSize), writers.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}

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
	if len(c.spec.ServiceAccountKeyJSON) != 0 {
		opts = append(opts, option.WithCredentialsJSON([]byte(c.spec.ServiceAccountKeyJSON)))
	}
	client, err := bigquery.NewClient(ctx, c.spec.ProjectID, opts...)
	if err != nil {
		return nil, err
	}
	if c.spec.DatasetLocation != "" {
		client.Location = c.spec.DatasetLocation
	}
	return client, nil
}

func (c *Client) Close(_ context.Context) error {
	return c.client.Close()
}
