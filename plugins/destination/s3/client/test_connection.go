package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/s3/v7/client/spec"
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
		{
			s := &spec.Spec{}
			err := json.Unmarshal(specBytes, &s)
			if err != nil {
				return fmt.Errorf("failed to unmarshal s3 spec: %w", err)
			}
			if err := s.Validate(); err != nil {
				return err
			}
			s.SetDefaults()
			b := true
			s.TestWrite = &b // Force test write to be enabled for connection testing
			specBytes, err = json.Marshal(s)
			if err != nil {
				return fmt.Errorf("failed to marshal s3 spec: %w", err)
			}
		}

		_, err := createClientFn(ctx, logger, specBytes, plugin.NewClientOptions{
			InvocationID: uuid.NewString(),
		})
		if err == nil {
			return nil
		}

		if errors.Is(err, errTestWriteFailed) {
			return plugin.NewTestConnError(codeUnauthorized, err)
		}

		return plugin.NewTestConnError(codeInvalidSpec, err)
	}
}
