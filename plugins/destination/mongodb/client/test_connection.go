package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
	codeUnauthorized     = "UNAUTHORIZED"
	codeUnreachable      = "UNREACHABLE"
)

func NewConnectionTester(createClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		if err == nil {
			return nil
		}

		switch err.(type) {
		case topology.ConnectionError:
			var authErr *auth.Error
			if errors.As(err, &authErr) {
				return plugin.NewTestConnError(codeUnauthorized, err)
			}
			return plugin.NewTestConnError(codeConnectionFailed, err)
		case topology.ServerSelectionError:
			return plugin.NewTestConnError(codeUnreachable, err)
		default:
			if errors.Is(err, errInvalidSpec) {
				return plugin.NewTestConnError(codeInvalidSpec, err)
			}
			return plugin.NewTestConnError(codeConnectionFailed, err)
		}
	}
}
