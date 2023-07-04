package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"

	"github.com/cloudquery/filetypes/v4"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	streamingbatchwriter.IgnoreMigrateTable
	streamingbatchwriter.UnimplementedDeleteStale

	logger zerolog.Logger
	spec   *Spec
	*filetypes.Client
	writer *streamingbatchwriter.StreamingBatchWriter

	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "s3").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal s3 spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	cfg.Region = c.spec.Region
	cfg.EndpointResolverWithOptions = c

	c.s3Client = s3.NewFromConfig(cfg)
	c.uploader = manager.NewUploader(c.s3Client)
	c.downloader = manager.NewDownloader(c.s3Client)

	if *c.spec.TestWrite {
		// we want to run this test because we want it to fail early if the bucket is not accessible
		timeNow := time.Now().UTC()
		if _, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(c.spec.Bucket),
			Key:    aws.String(replacePathVariables(c.spec.Path, "TEST_TABLE", "TEST_UUID", c.spec.Format, timeNow)),
			Body:   bytes.NewReader([]byte("")),
		}); err != nil {
			return nil, fmt.Errorf("failed to write test file to S3: %w", err)
		}
	}

	c.writer, err = streamingbatchwriter.New(c,
		streamingbatchwriter.WithBatchSizeRows(*c.spec.BatchSize),
		streamingbatchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
		streamingbatchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.writer.Close(ctx)
}

func (c *Client) ResolveEndpoint(service, region string, options ...any) (aws.Endpoint, error) {
	if c.spec.Endpoint == "" || service != s3.ServiceID {
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}

	return aws.Endpoint{
		PartitionID:   "aws",
		URL:           c.spec.Endpoint,
		SigningRegion: region,
		Source:        aws.EndpointSourceCustom,
	}, nil
}
