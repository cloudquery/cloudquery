package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
)

func NewConnectionTester(createClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		var tcErr *plugin.TestConnError
		if errors.As(err, &tcErr) {
			return tcErr
		}
		return err
	}
}
