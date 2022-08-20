package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func GcpMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, p *plugins.SourcePlugin, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to creattService %w", err)
		}
		var gcpSpec Spec
		if err := spec.UnmarshalSpec(&gcpSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		c := &Client{
			plugin: p,
			// logger:   t.Log(),
			Services: svc,
			projects: []string{"testProject"},
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		"dev",
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:   "dev",
		Tables: []string{"*"},
	})
}
