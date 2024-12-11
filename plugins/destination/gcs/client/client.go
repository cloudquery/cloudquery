package client

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/cloudquery/cloudquery/plugins/destination/gcs/v5/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"
	"google.golang.org/api/option"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	streamingbatchwriter.IgnoreMigrateTable
	streamingbatchwriter.UnimplementedDeleteStale
	streamingbatchwriter.UnimplementedDeleteRecords

	syncID string
	logger zerolog.Logger
	spec   *spec.Spec

	gcsClient *storage.Client
	bucket    *storage.BucketHandle
	*filetypes.Client

	writer *streamingbatchwriter.StreamingBatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, s []byte, newClientOpts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "gcs").Logger(),
		syncID: newClientOpts.InvocationID,
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcs spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(&c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	opts := []option.ClientOption{}
	if len(c.spec.ServiceAccountKeyJSON) != 0 {
		opts = append(opts, option.WithCredentialsJSON([]byte(c.spec.ServiceAccountKeyJSON)))
	}
	c.gcsClient, err = storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCP storage client: %w", err)
	}
	c.bucket = c.gcsClient.Bucket(c.spec.Bucket)
	// we upload it because we want to fail early if we don't have permissions
	gcpWriter := c.bucket.Object("/tmp/.cq-test-file-" + uuid.NewString()).NewWriter(ctx)
	if _, err := gcpWriter.Write([]byte("")); err != nil {
		return nil, fmt.Errorf("failed to write test file to GCS: %w", err)
	}
	if err := gcpWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close GCS writer: %w", err)
	}

	c.writer, err = streamingbatchwriter.New(c,
		streamingbatchwriter.WithBatchSizeRows(*c.spec.BatchSize),
		streamingbatchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
		streamingbatchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
		streamingbatchwriter.WithLogger(c.logger),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.writer.Close(ctx)
}
