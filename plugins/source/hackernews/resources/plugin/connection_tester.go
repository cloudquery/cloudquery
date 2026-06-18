package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/services"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	codeInvalidSpec      = "INVALID_SPEC"
	codeConnectionFailed = "CONNECTION_FAILED"
)

func connectionTester(newClient func() services.HackernewsClient) plugin.ConnectionTester {
	return func(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
		var s client.Spec
		if err := json.Unmarshal(specBytes, &s); err != nil {
			return plugin.NewTestConnError(codeInvalidSpec, fmt.Errorf("failed to unmarshal spec: %w", err))
		}
		s.SetDefaults()
		if err := s.Validate(); err != nil {
			return plugin.NewTestConnError(codeInvalidSpec, fmt.Errorf("failed to validate spec: %w", err))
		}

		if _, err := newClient().MaxItemID(ctx); err != nil {
			return plugin.NewTestConnError(codeConnectionFailed, err)
		}
		return nil
	}
}

func defaultClientFn() services.HackernewsClient {
	return hackernews.NewClient()
}
