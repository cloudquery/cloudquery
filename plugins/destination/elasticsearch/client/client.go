package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	logger      zerolog.Logger
	spec        *Spec
	client      *elasticsearch.Client
	typedClient *elasticsearch.TypedClient
	writer      *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "elasticsearch-dest").Logger(),
		spec:   &Spec{},
	}
	if err := json.Unmarshal(specBytes, c.spec); err != nil {
		return nil, err
	}

	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}
	retryBackoff := backoff.NewExponentialBackOff()
	var caCert []byte
	if len(c.spec.CACert) > 0 {
		caCert = []byte(c.spec.CACert)
	}
	cfg := elasticsearch.Config{
		Addresses:              c.spec.Addresses,
		Username:               c.spec.Username,
		Password:               c.spec.Password,
		CloudID:                c.spec.CloudID,
		APIKey:                 c.spec.APIKey,
		ServiceToken:           c.spec.ServiceToken,
		CertificateFingerprint: c.spec.CertificateFingerprint,
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

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		return fmt.Errorf("failed to close batch writer: %w", err)
	}
	return nil
}

func (*Client) getIndexNamePattern(table *schema.Table) string {
	hasPrimaryKeys := len(table.PrimaryKeys()) > 0
	if hasPrimaryKeys {
		return table.Name
	}
	return table.Name + "-*"
}

func (*Client) getIndexName(table *schema.Table, t time.Time) string {
	hasPrimaryKeys := len(table.PrimaryKeys()) > 0
	if hasPrimaryKeys {
		return table.Name
	}
	return table.Name + "-" + t.Format("2006-01-02")
}
