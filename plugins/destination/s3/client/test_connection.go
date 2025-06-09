package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type NewClientFn func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error)

const (
	codeInvalidSpec  string = "INVALID_SPEC"
	codeUnauthorized string = "UNAUTHORIZED"
)

func NewConnectionTester(createClientFn NewClientFn) plugin.ConnectionTester {
	return func(ctx context.Context, logger zerolog.Logger, specBytes []byte) error {
		invocationID, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		_, err = createClientFn(ctx, logger, specBytes, plugin.NewClientOptions{InvocationID: invocationID.String()})
		if err == nil {
			return nil
		}

		if errors.Is(err, errTestWriteFailed) {
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeInvalidSpec, err)
	}
}
