package client

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	*filetypes.Client
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "s3").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal s3 spec: %w", err)
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

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	location, err := getBucketLocation(ctx, s3.NewFromConfig(cfg), c.pluginSpec.Bucket)
	if err != nil {
		return nil, fmt.Errorf("unable to determine region of S3 bucket: %w", err)
	}
	cfg.Region = location
	c.s3Client = s3.NewFromConfig(cfg)
	c.uploader = manager.NewUploader(c.s3Client)
	c.downloader = manager.NewDownloader(c.s3Client)

	// we want to run this test because we want it to fail early if the bucket is not accessible
	timeNow := time.Now().UTC()
	if _, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.pluginSpec.Bucket),
		Key:    aws.String(replacePathVariables(spec.Path, "TEST_TABLE", "TEST_UUID", timeNow)),
		Body:   bytes.NewReader([]byte("")),
	}); err != nil {
		return nil, fmt.Errorf("failed to write test file to S3: %w", err)
	}
	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func getBucketLocation(ctx context.Context, s3Client *s3.Client, bucket string) (string, error) {
	output, err := s3Client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return "", fmt.Errorf("failed to get bucket location: %w", err)
	}
	if output.LocationConstraint == "" {
		return "us-east-1", nil
	}
	return string(output.LocationConstraint), nil
}
