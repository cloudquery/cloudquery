package client

import (
	"context"
	"errors"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
	codeUnauthorized     = "UNAUTHORIZED"
	codeUnreachable      = "UNREACHABLE"
)

const (
	errUnauthorized = "an error happened during the Info query execution: EOF"
	errUnreachable  = "an error happened during the Info query execution: dial tcp"
)

func NewConnectionTester(createClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		if err == nil {
			return nil
		}

		if errors.Is(err, errInvalidSpec) {
			return plugin.NewTestConnError(codeInvalidSpec, err)
		}

		if strings.Contains(err.Error(), errUnreachable) {
			return plugin.NewTestConnError(codeUnreachable, err)
		}
		if strings.Contains(err.Error(), errUnauthorized) {
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeConnectionFailed, err)
	}
}
