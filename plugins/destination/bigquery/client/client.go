package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/bigquery"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/bigquery/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

const cloudQueryGPN = "CloudQuery"

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   Spec
	client *bigquery.Client
	writer *batchwriter.BatchWriter

	batchwriter.UnimplementedDeleteStale
	batchwriter.UnimplementedDeleteRecord
}

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "bq-dest").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}
	if err := json.Unmarshal(specBytes, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BigQuery spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = batchwriter.New(c,
		batchwriter.WithLogger(c.logger),
		batchwriter.WithBatchSize(c.spec.BatchSize),
		batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes),
		batchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
	)
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

	if err := c.validateCreds(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to validate credentials: %w", err)
	}

	return c, nil
}

func (c *Client) bqClient(ctx context.Context) (*bigquery.Client, error) {
	opts := []option.ClientOption{
		option.WithRequestReason("CloudQuery BigQuery destination"),
		option.WithUserAgent(fmt.Sprintf("CloudQuery_BigQuery_Destination/%s (GPN:%s)", internalPlugin.Version, cloudQueryGPN)),
	}
	if len(c.spec.ServiceAccountKeyJSON) != 0 {
		opts = append(opts, option.WithCredentialsJSON([]byte(c.spec.ServiceAccountKeyJSON)))
	}
	if c.spec.Endpoint != "" {
		opts = append(opts, option.WithEndpoint(c.spec.Endpoint))
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

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		return err
	}
	return c.client.Close()
}

func (c *Client) validateCreds(ctx context.Context) error {
	datasetRef := c.client.Dataset(c.spec.DatasetID)
	_, err := datasetRef.Metadata(ctx)
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				return fmt.Errorf("invalid dataset. dataset must be created before sync or migration: %w", err)
			}
		}
		return fmt.Errorf("failed to validate credentials: %w", err)
	}
	return nil
}
