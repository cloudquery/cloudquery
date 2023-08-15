package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"

	"github.com/rs/zerolog"
)

type Client struct {
	firehoseClient *firehose.Client
	spec           Spec

	logger zerolog.Logger
	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	parsedARN, err := arn.Parse(spec.StreamARN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse firehose stream ARN: %w", err)
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(parsedARN.Region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	return &Client{
		logger:         logger.With().Str("module", "firehose").Logger(),
		spec:           spec,
		firehoseClient: firehose.NewFromConfig(cfg),
	}, nil
}

func (*Client) Close(context.Context) error { return nil }

func (*Client) Read(context.Context, *schema.Table, chan<- arrow.Record) error {
	return plugin.ErrNotImplemented
}
