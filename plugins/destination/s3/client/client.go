package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/plugins/destination/s3/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"

	"github.com/cloudquery/filetypes/v4"
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
	*filetypes.Client
	writer *streamingbatchwriter.StreamingBatchWriter

	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
}

func New(ctx context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "s3").Logger(),
		syncID: opts.InvocationID,
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal s3 spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	if c.syncID == "" && c.spec.PathContainsSyncID() {
		return nil, fmt.Errorf("path contains {{SYNC_ID}}. Upgrade your CLI to use this path variable")
	}

	filetypesClient, err := filetypes.NewClient(&c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	cfg.Region = c.spec.Region

	cfg.HTTPClient = awshttp.NewBuildableClient().WithTransportOptions(func(tr *http.Transport) {
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}
		tr.TLSClientConfig.InsecureSkipVerify = c.spec.EndpointSkipTLSVerify
	})
	c.s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		if len(c.spec.Endpoint) > 0 {
			baseEndpoint := c.spec.Endpoint
			o.BaseEndpoint = &baseEndpoint
		}
		o.UsePathStyle = c.spec.UsePathStyle
	})
	c.uploader = manager.NewUploader(c.s3Client)
	c.downloader = manager.NewDownloader(c.s3Client)

	if *c.spec.TestWrite {
		// we want to run this test because we want it to fail early if the bucket is not accessible
		timeNow := time.Now().UTC()

		params := &s3.PutObjectInput{
			Bucket: aws.String(c.spec.Bucket),
			Key:    aws.String(c.spec.ReplacePathVariables("TEST_TABLE", "TEST_UUID", timeNow, c.syncID)),
			Body:   bytes.NewReader([]byte("")),
		}

		sseConfiguration := c.spec.ServerSideEncryptionConfiguration
		if sseConfiguration != nil {
			params.SSEKMSKeyId = &sseConfiguration.SSEKMSKeyId
			params.ServerSideEncryption = sseConfiguration.ServerSideEncryption
		}

		if _, err := c.uploader.Upload(ctx, params); err != nil {
			return nil, fmt.Errorf("failed to write test file to S3: %w", err)
		}
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
