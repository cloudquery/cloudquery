package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/xkcd/client"
	"github.com/cloudquery/cloudquery/plugins/source/xkcd/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
)

func connectionTester(newClient func() (*xkcd.Client, error)) plugin.ConnectionTester {
	return func(_ context.Context, _ zerolog.Logger, specBytes []byte) error {
		var s client.Spec
		if err := json.Unmarshal(specBytes, &s); err != nil {
			return plugin.NewTestConnError(codeInvalidSpec, fmt.Errorf("failed to unmarshal spec: %w", err))
		}
		s.SetDefaults()
		if err := s.Validate(); err != nil {
			return plugin.NewTestConnError(codeInvalidSpec, fmt.Errorf("failed to validate spec: %w", err))
		}

		c, err := newClient()
		if err != nil {
			return plugin.NewTestConnError(codeConnectionFailed, fmt.Errorf("failed to create client: %w", err))
		}

		if _, err := c.GetLatestComic(0); err != nil {
			return plugin.NewTestConnError(codeConnectionFailed, err)
		}
		return nil
	}
}

func defaultClientFn() (*xkcd.Client, error) {
	return xkcd.NewClient()
}
