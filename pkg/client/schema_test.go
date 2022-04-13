package client

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
)

func Test_GetProviderSchema(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	pManager, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL,
		registry.WithPluginDirectory(filepath.Join(".", ".cq", "providers"))), plugin.WithAllowReattach())
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	s, err := GetProviderSchema(context.Background(), pManager, &GetProviderSchemaOptions{Provider: registry.Provider{
		Name:    "test",
		Version: "latest",
		Source:  registry.DefaultOrganization,
	}})
	if s == nil {
		t.FailNow()
	}
	assert.Equal(t, "test", s.Name)
	assert.Equal(t, "v0.0.0", s.Version)
	assert.Equal(t, 3, len(s.ResourceTables))
	assert.Nil(t, err)
}
