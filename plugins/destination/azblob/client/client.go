package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	storageClient *azblob.Client

	*filetypes.Client
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var err error
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "azb").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal azblob spec: %w", err)
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
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure credential: %w", err)
	}
	c.storageClient, err = azblob.NewClient("https://"+c.pluginSpec.StorageAccount+".blob.core.windows.net", cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure storage client: %w", err)
	}

	_, err = c.storageClient.UploadStream(ctx, c.pluginSpec.Container, "cq-test-file", strings.NewReader(""), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to write test file to Azure: %w", err)
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
