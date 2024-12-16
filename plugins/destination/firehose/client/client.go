package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"

	"github.com/rs/zerolog"
)

type Client struct {
	firehoseClient *firehose.Client
	spec           spec.Spec

	logger zerolog.Logger
	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var s spec.Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return nil, err
	}

	parsedARN, err := arn.Parse(s.StreamARN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse firehose stream ARN: %w", err)
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(parsedARN.Region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}
	if err := validateCredentials(ctx, cfg); err != nil {
		return nil, err
	}
	return &Client{
		logger:         logger.With().Str("module", "firehose").Logger(),
		spec:           s,
		firehoseClient: firehose.NewFromConfig(cfg),
	}, nil
}

func (*Client) Close(context.Context) error { return nil }

func (*Client) Read(context.Context, *schema.Table, chan<- arrow.Record) error {
	return plugin.ErrNotImplemented
}

func validateCredentials(ctx context.Context, cfg aws.Config) error {
	stsClient := sts.NewFromConfig(cfg)
	_, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return fmt.Errorf("failed to validate AWS credentials: %w", err)
	}
	return err
}
