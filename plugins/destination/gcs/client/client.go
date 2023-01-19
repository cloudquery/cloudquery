package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	gcsClient *storage.Client
	bucket    *storage.BucketHandle
	*filetypes.Client
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var err error
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "gcs").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(c.pluginSpec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	c.gcsClient, err = storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCP storage client: %w", err)
	}
	c.bucket = c.gcsClient.Bucket(c.pluginSpec.Bucket)
	// we upload it because we want to fail early if we don't have permissions
	gcpWriter := c.bucket.Object("/tmp/.cq-test-file-" + uuid.NewString()).NewWriter(ctx)
	if _, err := gcpWriter.Write([]byte("")); err != nil {
		return nil, fmt.Errorf("failed to write test file to GCS: %w", err)
	}
	if err := gcpWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close GCS writer: %w", err)
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
