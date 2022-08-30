package core

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/cli/internal/firebase"
	"github.com/cloudquery/cloudquery/cli/pkg/config"
	"github.com/cloudquery/cloudquery/cli/pkg/plugin"
	"github.com/cloudquery/cloudquery/cli/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadExisting(t *testing.T) {
	tempDir := t.TempDir()
	pm, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
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
	require.Nil(t, diags)

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
	pm, err = plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
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

// TODO: latest + unverified provider won't work on download
func TestDownloadCommunity(t *testing.T) {
	tempDir := t.TempDir()
	pm, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	src, name, err := ParseProviderSource(&config.RequiredProvider{
		Name:    "yandex-cloud/yandex",
		Source:  nil,
		Version: "v0.0.8",
	})
	assert.Nil(t, err)
	// Download plugin for the first time
	_, diags := Download(context.Background(), pm, &DownloadOptions{
		Providers: []registry.Provider{
			{
				Name:    name,
				Version: "v0.0.8",
				Source:  src,
			},
		},
		NoVerify: false,
	})
	assert.Nil(t, diags)

	source := "yandex-cloud"
	src, name, err = ParseProviderSource(&config.RequiredProvider{
		Name:    "yandex",
		Source:  &source,
		Version: "v0.0.8",
	})
	assert.Nil(t, err)
	// Download plugin for the first time
	_, diags = Download(context.Background(), pm, &DownloadOptions{
		Providers: []registry.Provider{
			{
				Name:    name,
				Version: "v0.0.8",
				Source:  src,
			},
		},
		NoVerify: false,
	})
	assert.Nil(t, diags)
}
