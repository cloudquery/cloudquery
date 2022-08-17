package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func GcpMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false
	// 	cfg := `
	// tables: ["*"]
	// `
	// plugins.TestResource(t, plugins.ResourceTestCase{
	// 	Plugin: &plugins.SourcePlugin{
	// 		Name:    "gcp_mock_test_provider",
	// 		Version: "development",
	// 		Configure: func(ctx context.Context, p *plugins.SourcePlugin, s specs.SourceSpec) (schema.ClientMeta, error) {
	// 			svc, err := createService()
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			var gcpSpec Spec
	// 			if err := s.Spec.Decode(&gcpSpec); err != nil {
	// 				return nil, fmt.Errorf("failed to decode gcp spec: %w", err)
	// 			}
	// 			c := &Client{
	// 				plugin:   p,
	// 				logger:   p.Logger,
	// 				Services: svc,
	// 				projects: []string{"testProject"},
	// 			}

	// 			return c, nil
	// 		},
	// 		Tables: []*schema.Table{
	// 			table,
	// 		},
	// 		// Config: func() provider.Config {
	// 		// 	return &Config{}
	// 		// },
	// 	},
	// 	Config: cfg,
	// })
}
