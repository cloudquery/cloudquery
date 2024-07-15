package client

import (
	"context"
	"errors"
	"net"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type NewClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
	codeUnauthorized     = "UNAUTHORIZED"
	codeUnreachable      = "UNREACHABLE"
)

func NewConnectionTester(createClientFn NewClientFn) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		if err == nil {
			return nil
		}

		if errors.Is(err, errInvalidSpec) {
			return plugin.NewTestConnError(codeInvalidSpec, err)
		}

		var netErr *net.OpError
		if errors.As(err, &netErr) {
			return plugin.NewTestConnError(codeUnreachable, err)
		}

		var mssqlErr interface{ SQLErrorNumber() int32 }
		if errors.As(err, &mssqlErr) {
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeConnectionFailed, err)
	}
}
