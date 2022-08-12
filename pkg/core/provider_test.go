package core

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const expectedProviderConfig = `configuration:

#account:
#  name: "1"
#  id: testid
#  regions:
#    - asdas
#  
# list of resources to fetch
resources:
    - error_resource
    - migrate_resource
    - panic_resource
    - slow_resource
    - very_slow_resource
`

type mockVersionsClient struct {
}

func (mvc mockVersionsClient) GetLatestProviderRelease(ctx context.Context, org, pluginType, pluginName string) (string, error) {
	return "v0.1.21", nil // latest version of test provider before it was archived
}

func Test_CheckAvailableUpdates(t *testing.T) {
	latestVersion := getLatestVersion(t, "test")

	testCases := []struct {
		Name    string
		Options *CheckUpdatesOptions

		ExpectedAvailableUpdates []AvailableUpdate
		ExpectedDiags            diag.FlatDiags
	}{
		{
			Name: "simple-update-check",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "test", Version: "v0.0.1", Source: registry.DefaultOrganization},
			}},
			ExpectedAvailableUpdates: []AvailableUpdate{{
				Name:             "test",
				CurrentVersion:   "v0.0.1",
				AvailableVersion: latestVersion,
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
			ExpectedDiags: diag.FlatDiags{
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
		{
			Name: "check-up-to-date",
			Options: &CheckUpdatesOptions{Providers: []registry.Provider{
				{Name: "test", Version: latestVersion, Source: registry.DefaultOrganization},
			}},
			ExpectedAvailableUpdates: []AvailableUpdate{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			reg := registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(t.TempDir()))
			updates, diags := CheckAvailableUpdates(context.Background(), reg, tc.Options)
			if len(tc.ExpectedDiags) > 0 {
				assert.Equal(t, tc.ExpectedDiags, diag.FlattenDiags(diags, true))
			} else {
				assert.Len(t, tc.ExpectedDiags, 0)
			}
			if tc.ExpectedAvailableUpdates != nil {
				assert.Equal(t, tc.ExpectedAvailableUpdates, updates)
			}
		})
	}
}

func Test_GetProviderConfig(t *testing.T) {
	provider := registry.Provider{
		Name:    "test",
		Source:  "cloudquery",
		Version: "v0.1.17",
	}
	pm, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL))
	assert.Nil(t, err)
	_, diags := Download(context.TODO(), pm, &DownloadOptions{
		Providers: []registry.Provider{provider},
		NoVerify:  false,
	})
	assert.False(t, diags.HasErrors())
	defer pm.Shutdown()

	ctx := context.Background()
	pConfig, diags := GetProviderConfiguration(ctx, pm, &GetProviderConfigOptions{
		Provider: provider,
	})
	if diags.HasErrors() {
		t.FailNow()
	}
	assert.NotNil(t, pConfig)
	assert.Equal(t, expectedProviderConfig, string(pConfig.Config))

	var yCfg yaml.Node
	err = yaml.Unmarshal(pConfig.Config, &yCfg)
	assert.NoError(t, err)
}

func getLatestVersion(t *testing.T, name string) string {
	mvc := mockVersionsClient{}
	reg := registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(t.TempDir()), registry.WithVersionsClient(mvc))
	latest, diags := CheckAvailableUpdates(context.Background(), reg, &CheckUpdatesOptions{Providers: []registry.Provider{
		{Name: name, Version: "v0.0.0", Source: registry.DefaultOrganization},
	}})
	assert.Nil(t, diags)
	assert.NotNil(t, latest)
	assert.Len(t, latest, 1)
	// get latest version
	return latest[0].AvailableVersion
}
