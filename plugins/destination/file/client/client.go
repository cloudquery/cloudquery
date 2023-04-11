package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	*filetypes.Client
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

	filetypesClient, err := filetypes.NewClient(c.pluginSpec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

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
