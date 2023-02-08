package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger      zerolog.Logger
	spec        specs.Destination
	metrics     destination.Metrics
	pluginSpec  Spec
	client      *elasticsearch.Client
	typedClient *elasticsearch.TypedClient
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "elasticsearch-dest").Logger(),
		spec:   destSpec,
	}
	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Elasticsearch spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec
	retryBackoff := backoff.NewExponentialBackOff()
	var caCert []byte
	if len(spec.CACert) > 0 {
		caCert = []byte(spec.CACert)
	}
	cfg := elasticsearch.Config{
		Addresses:              spec.Addresses,
		Username:               spec.Username,
		Password:               spec.Password,
		CloudID:                spec.CloudID,
		APIKey:                 spec.APIKey,
		ServiceToken:           spec.ServiceToken,
		CertificateFingerprint: spec.CertificateFingerprint,
		CACert:                 caCert,
		// Retry on 429 TooManyRequests statuses
		RetryOnStatus: []int{502, 503, 504, 429},
		// Configure the backoff function
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		// Retry up to 5 attempts
		MaxRetries: 5,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}
	info, err := es.Info().Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Elasticsearch cluster info: %w", err)
	}
	defer info.Body.Close()
	b, err := io.ReadAll(info.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Elasticsearch cluster info response: %w", err)
	}
	c.logger.Debug().Str("cluster_info", string(b)).Msg("Elasticsearch cluster info")
	c.typedClient = es
	c.client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create untyped Elasticsearch client: %w", err)
	}
	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) getIndexNamePattern(tableName string) string {
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return tableName + "-*"
	case specs.WriteModeOverwrite:
		return tableName
	case specs.WriteModeOverwriteDeleteStale:
		return tableName
	default:
		return ""
	}
}

func (c *Client) getIndexName(tableName string, t time.Time) string {
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return tableName + "-" + t.Format("2006-01-02")
	case specs.WriteModeOverwrite:
		return tableName
	case specs.WriteModeOverwriteDeleteStale:
		return tableName
	default:
		return ""
	}
}
