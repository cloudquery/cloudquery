package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func K8sMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	mockController := gomock.NewController(t)

	configureFunc := func(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
		var k8sSpec Spec
		if err := s.UnmarshalSpec(&k8sSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal k8s spec: %w", err)
		}

		c := &Client{
			logger:  logger,
			Context: "testContext",
			spec:    &k8sSpec,
		}
		c.SetServices(map[string]Services{"testContext": builder(t, mockController)})
		return c, nil
	}

	plugin := plugins.NewSourcePlugin(
		table.Name,
		"dev",
		[]*schema.Table{
			table,
		},
		configureFunc,
	)

	plugins.TestSourcePluginSync(t, plugin, specs.Source{
		Name:   "dev",
		Tables: []string{table.Name},
	})
}
