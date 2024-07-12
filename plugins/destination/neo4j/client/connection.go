package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type NewClientFunc func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)

const (
	codeInvalidSpec      string = "INVALID_SPEC"
	codeUnreachable      string = "UNREACHABLE"
	codeUnauthorized     string = "UNAUTHORIZED"
	codeConnectionFailed string = "CONNECTION_FAILED"
)

var (
	errInvalidSpec  = errors.New("invalid spec")
	errUnreachable  = errors.New("unreachable")
	errUnauthorized = errors.New("unauthorized")
)

func NewConnectionTester(createClientFunc NewClientFunc) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, specBytes []byte) (err error) {
		if _, err = createClientFunc(ctx, logger, specBytes, plugin.NewClientOptions{}); err == nil {
			return nil
		}

		switch {
		case errors.Is(err, errInvalidSpec):
			return plugin.NewTestConnError(codeInvalidSpec, err)

		case errors.Is(err, errUnreachable):
			return plugin.NewTestConnError(codeUnreachable, err)

		case errors.Is(err, errUnauthorized):
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeConnectionFailed, err)
	}
}
