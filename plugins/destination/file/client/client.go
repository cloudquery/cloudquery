package client

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/filetypes/parquet"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	formatClient formatClient

	// Embedded transformers
	schema.CQTypeTransformer
	reverseTransformer
}

type reverseTransformer interface {
	ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error)
}

type formatClient interface {
	Read(r io.Reader, table *schema.Table, sourceName string, res chan<- []any) error
	WriteTableBatch(w io.Writer, table *schema.Table, resources [][]any) error
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("file destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "file").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()

	var err error
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		c.formatClient, err = csv.NewClient()
		c.CQTypeTransformer = &csv.Transformer{}
		c.reverseTransformer = &csv.ReverseTransformer{}
	case FormatTypeJSON:
		c.formatClient, err = json.NewClient()
		c.CQTypeTransformer = &json.Transformer{}
		c.reverseTransformer = &json.ReverseTransformer{}
	case FormatTypeParquet:
		c.formatClient, err = parquet.NewClient()
		c.CQTypeTransformer = &parquet.Transformer{}
		c.reverseTransformer = &parquet.ReverseTransformer{}
	default:
		return nil, fmt.Errorf("unknown format %q", c.pluginSpec.Format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create filetype client: %w", err)
	}

	if err := os.MkdirAll(c.pluginSpec.Directory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
