package client

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	plugins.DefaultReverseTransformer
	logger  zerolog.Logger
	spec    specs.Destination
	csvSpec Spec

	awsUploader      *manager.Uploader
	gcpStorageClient *storage.Client

	metrics plugins.DestinationMetrics
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("file destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "dest-file").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.csvSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	c.csvSpec.SetDefaults()

	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		if err := os.MkdirAll(c.csvSpec.Directory, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	case BackendTypeS3:
		awsCfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS SDK config, %w", err)
		}
		awsClient := s3.NewFromConfig(awsCfg)
		c.awsUploader = manager.NewUploader(awsClient)
		if _, err := c.awsUploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(c.csvSpec.Directory),
			Key:    aws.String("cq-test-file"),
			Body:   bytes.NewReader([]byte("test")),
		}); err != nil {
			return nil, fmt.Errorf("failed to write test file to S3: %w", err)
		}
	case BackendTypeGCS:
		var err error
		c.gcpStorageClient, err = storage.NewClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create GCP storage client %w", err)
		}
		if _, err := c.gcpStorageClient.Bucket(c.csvSpec.Directory).Object("cq-test-file").NewWriter(ctx).Write([]byte("test")); err != nil {
			return nil, fmt.Errorf("failed to write test file to gs://%s %w", c.csvSpec.Directory, err)
		}
	default:
		return nil, fmt.Errorf("unknown backend: %s", c.csvSpec.Backend)
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}
