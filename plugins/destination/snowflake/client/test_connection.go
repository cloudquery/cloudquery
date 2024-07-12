package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/snowflakedb/gosnowflake"
)

type NewClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)

const (
	codeConnectionFailed string = "CONNECTION_FAILED"
	codeInvalidSpec      string = "INVALID_SPEC"
	codeUnauthorized     string = "UNAUTHORIZED"
	codeUnreachable      string = "UNREACHABLE"
)

func NewConnectionTester(createClientFn NewClientFn) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, spec []byte) error {
		_, err := createClientFn(ctx, logger, spec, plugin.NewClientOptions{})
		if err == nil {
			return nil
		}

		var snowflakeErr *gosnowflake.SnowflakeError
		if errors.As(err, &snowflakeErr) {
			switch snowflakeErr.Number {
			case gosnowflake.ErrCodeFailedToConnect, gosnowflake.ErrCodeServiceUnavailable, gosnowflake.ErrCodeIdpConnectionError:
				return plugin.NewTestConnError(codeUnreachable, err)
			case gosnowflake.ErrFailedToAuth, gosnowflake.ErrFailedToAuthSAML, gosnowflake.ErrFailedToAuthOKTA, gosnowflake.ErrObjectNotExistOrAuthorized:
				return plugin.NewTestConnError(codeUnauthorized, err)
			default:
				return plugin.NewTestConnError(codeConnectionFailed, err)
			}
		}

		if errors.Is(err, errInvalidSpec) {
			return plugin.NewTestConnError(codeInvalidSpec, err)
		}

		return plugin.NewTestConnError(codeConnectionFailed, err)
	}
}
