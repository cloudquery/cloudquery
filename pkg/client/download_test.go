package client

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
)

func TestDownloadExisting(t *testing.T) {
	tempDir := t.TempDir()
	pm, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	// Plugin shouldn't exist
	_, err = pm.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    "test",
			Version: "latest",
			Source:  "cloudquery",
		},
		Alias: "",
		Env:   nil,
	})
	assert.Error(t, err)

	// Download plugin for the first time
	_, diags := Download(context.Background(), pm, &DownloadOptions{
		Providers: []registry.Provider{
			{
				Name:    "test",
				Version: "v0.0.11",
				Source:  "cloudquery",
			},
		},
		NoVerify: false,
	})
	assert.Nil(t, diags)

	// Plugin is downloaded and should be created
	p, err := pm.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    "test",
			Version: "latest",
			Source:  "cloudquery",
		},
		Alias: "",
		Env:   nil,
	})
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), "test")
	assert.Equal(t, p.Version(), "v0.0.11")
	pm.ClosePlugin(p)

	// Create a new clean plugin.Manager, and check that it shouldn't download again
	pm, err = plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	// Plugin is downloaded and should be created
	p, err = pm.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    "test",
			Version: "latest",
			Source:  "cloudquery",
		},
		Alias: "",
		Env:   nil,
	})
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), "test")
	assert.Equal(t, p.Version(), "v0.0.11")
	pm.ClosePlugin(p)
}

func TestDownloadUnverified(t *testing.T) {
	tempDir := t.TempDir()
	pm, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	// Plugin shouldn't exist
	_, err = pm.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    "unverified",
			Version: "latest",
			Source:  "cloudquery",
		},
		Alias: "",
		Env:   nil,
	})
	assert.Error(t, err)
	// Download plugin for the first time
	_, diags := Download(context.Background(), pm, &DownloadOptions{
		Providers: []registry.Provider{
			{
				Name:    "unverified",
				Version: "v0.0.11",
				Source:  "cloudquery",
			},
		},
		NoVerify: false,
	})
	assert.NotNil(t, diags)
}
