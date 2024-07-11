package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

const (
	codeInvalidSpec  = "INVALID_SPEC"
	codeUnauthorized = "UNAUTHORIZED"
)

func NewConnectionTester(createClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		if err == nil {
			return nil
		}
		if errors.Is(err, errValidateConnectionFailed) {
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeInvalidSpec, err)
	}
}
