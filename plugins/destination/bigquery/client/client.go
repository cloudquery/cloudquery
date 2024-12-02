package client

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/bigquery/v4/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"
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

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
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
	c.client, err = bqClient(context.Background(), c.spec)
	if err != nil {
		return nil, err
	}

	if err := validateCreds(ctx, c.client, c.spec.DatasetID); err != nil {
		return nil, fmt.Errorf("failed to validate credentials: %w", err)
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		return err
	}
	return c.client.Close()
}

func bqClient(ctx context.Context, s Spec) (*bigquery.Client, error) {
	opts := []option.ClientOption{
		option.WithRequestReason("CloudQuery BigQuery destination"),
		option.WithUserAgent(fmt.Sprintf("CloudQuery_BigQuery_Destination/%s (GPN:%s)", internalPlugin.Version, cloudQueryGPN)),
	}
	if len(s.ServiceAccountKeyJSON) != 0 {
		opts = append(opts, option.WithCredentialsJSON([]byte(s.ServiceAccountKeyJSON)))
	}
	if s.Endpoint != "" {
		opts = append(opts, option.WithEndpoint(s.Endpoint))
	}
	client, err := bigquery.NewClient(ctx, s.ProjectID, opts...)
	if err != nil {
		return nil, err
	}
	if s.DatasetLocation != "" {
		client.Location = s.DatasetLocation
	}
	return client, nil
}

func validateCreds(ctx context.Context, c *bigquery.Client, datasetID string) error {
	datasetRef := c.Dataset(datasetID)
	_, err := datasetRef.Metadata(ctx)
	if err != nil {
		if isAPINotFoundError(err) {
			return fmt.Errorf("invalid dataset. dataset must be created before sync or migration: %w", err)
		}
		return fmt.Errorf("failed to validate credentials: %w", err)
	}
	return nil
}

func TestConnection(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var s Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return &plugin.TestConnError{
			Code:    "INVALID_SPEC",
			Message: fmt.Errorf("failed to unmarshal spec: %w", err),
		}
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return err
	}

	c, err := bqClient(ctx, s)
	if err != nil {
		return err
	}

	if err := validateCreds(ctx, c, s.DatasetID); err != nil {
		return err
	}

	return c.Close()
}
