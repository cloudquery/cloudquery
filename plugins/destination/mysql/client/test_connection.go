package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

func NewConnectionTester(createClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		return err
	}
}
