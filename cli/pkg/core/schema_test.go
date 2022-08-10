package core

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/cli/internal/firebase"
	"github.com/cloudquery/cloudquery/cli/pkg/plugin"
	"github.com/cloudquery/cloudquery/cli/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetProviderSchema(t *testing.T) {
	provider := registry.Provider{
		Name:    "test",
		Source:  registry.DefaultOrganization,
		Version: "v0.0.11",
	}
	pm, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Nil(t, err)
	_, diags := Download(context.TODO(), pm, &DownloadOptions{
		Providers: []registry.Provider{provider},
		NoVerify:  false,
	})
	require.False(t, diags.HasDiags())

	s, err := GetProviderSchema(context.Background(), pm, &GetProviderSchemaOptions{Provider: provider})
	if s == nil {
		t.FailNow()
	}
	assert.Equal(t, "test", s.Name)
	assert.Equal(t, "v0.0.11", s.Version)
	assert.Equal(t, 5, len(s.ResourceTables))
	assert.Nil(t, err)
}
