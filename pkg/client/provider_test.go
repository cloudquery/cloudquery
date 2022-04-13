package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
)

func Test_CheckAvailableUpdates(t *testing.T) {
	testCases := []struct {
		Name    string
		Options *CheckUpdatesOptions

		ExpectedAvailableUpdates []AvailableUpdate
		ExpectedDiags            []diag.FlatDiag
	}{
		{
			Name: "simple-update-check",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "test", Version: "v0.0.1", Source: registry.DefaultOrganization},
			}},
			ExpectedAvailableUpdates: []AvailableUpdate{{
				Name:             "test",
				CurrentVersion:   "v0.0.1",
				AvailableVersion: "v0.0.11",
			}},
		},
		{
			Name: "check-update-with-latest-not-on-disk",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization},
			}},
			ExpectedAvailableUpdates: nil,
		},
		{
			Name: "check-no-existing",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "not-existing", Version: "v0.0.1", Source: registry.DefaultOrganization},
			}},
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:      "failed to find provider[not-existing] latest version",
					Type:     diag.INTERNAL,
					Severity: diag.ERROR,
					Summary:  "failed to find provider[not-existing] latest version",
				},
			},
			ExpectedAvailableUpdates: nil,
		},
		{
			Name: "check-with-higher-version",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "test", Version: "v999.999.999", Source: registry.DefaultOrganization},
			}},
			ExpectedAvailableUpdates: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			reg := registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(t.TempDir()))
			updates, diags := CheckAvailableUpdates(context.Background(), reg, tc.Options)
			if tc.ExpectedDiags != nil {
				assert.Equal(t, tc.ExpectedDiags, diag.FlattenDiags(diags, true))
			} else {
				assert.Len(t, tc.ExpectedDiags, 0)
			}
			assert.Equal(t, tc.ExpectedAvailableUpdates, updates)
		})
	}
}
