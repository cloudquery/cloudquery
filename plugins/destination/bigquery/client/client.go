package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	plugins.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	metrics    plugins.DestinationMetrics
	client     *bigquery.Client
	pluginSpec Spec
	projectID  string
	datasetID  string
}

func New(_ context.Context, logger zerolog.Logger, destSpec specs.Destination) (plugins.DestinationClient, error) {
	if destSpec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("bigquery destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "bq-dest").Logger(),
	}
	var spec Spec
	c.spec = destSpec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bigquery spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	client, err := bigquery.NewClient(context.Background(), spec.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create new BigQuery client: %w", err)
	}
	c.client = client
	c.projectID = spec.ProjectID
	c.datasetID = spec.DatasetID
	c.pluginSpec = spec
	return c, nil
}

func (c *Client) Close(_ context.Context) error {
	var err error
	if c.client == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	err = c.client.Close()
	c.client = nil
	return err
}
