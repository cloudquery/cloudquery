package client

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
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
	authFailedStr = "Authentication failed"
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

		var pe *proto.Exception
		if errors.As(err, &pe) {
			if strings.Contains(pe.Message, authFailedStr) {
				return plugin.NewTestConnError(codeUnauthorized, err)
			}
		}

		var opErr *net.OpError
		if errors.As(err, &opErr) {
			return plugin.NewTestConnError(codeUnreachable, err)
		}

		return plugin.NewTestConnError(codeConnectionFailed, err)
	}
}
