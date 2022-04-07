package client

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func Test_GetProviderSchema(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	pManager, err := plugin.NewManager(hclog.Default(), filepath.Join(".", ".cq", "providers"), registry.CloudQueryRegistryURL, nil)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	s, err := GetProviderSchema(context.Background(), pManager, &GetProviderSchemaOptions{"test"})
	if s == nil {
		t.FailNow()
	}
	assert.Equal(t, "test", s.Name)
	assert.Equal(t, "v0.0.0", s.Version)
	assert.Equal(t, 3, len(s.ResourceTables))
	assert.Nil(t, err)
}
