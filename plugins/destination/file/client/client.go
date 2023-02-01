package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/filetypes/parquet"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	CSVClient                 *csv.Client
	JSONClient                *json.Client
	ParquetClient             *parquet.Client
	csvTransformer            *csv.Transformer
	csvReverseTransformer     *csv.ReverseTransformer
	jsonTransformer           *json.Transformer
	jsonReverseTransformer    *json.ReverseTransformer
	parquetTransformer        *parquet.Transformer
	parquetReverseTransformer *parquet.ReverseTransformer
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("file destination only supports append mode")
	}
	c := &Client{
		logger:                    logger.With().Str("module", "file").Logger(),
		spec:                      spec,
		csvTransformer:            &csv.Transformer{},
		jsonTransformer:           &json.Transformer{},
		parquetTransformer:        &parquet.Transformer{},
		csvReverseTransformer:     &csv.ReverseTransformer{},
		jsonReverseTransformer:    &json.ReverseTransformer{},
		parquetReverseTransformer: &parquet.ReverseTransformer{},
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()

	csvClient, err := csv.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV client: %w", err)
	}
	c.CSVClient = csvClient

	jsonClient, err := json.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create JSON client: %w", err)
	}
	c.JSONClient = jsonClient

	parquetClient, err := parquet.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Parquet client: %w", err)
	}
	c.ParquetClient = parquetClient

	if err := os.MkdirAll(c.pluginSpec.Directory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
