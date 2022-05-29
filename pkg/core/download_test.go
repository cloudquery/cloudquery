package core

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
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

func TestDownloadUnverified(t *testing.T) {
	tempDir := t.TempDir()
	pm, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(tempDir)))
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
				Version: "v0.0.3",
				Source:  "cloudquery",
			},
		},
		NoVerify: false,
	})
	assert.NotNil(t, diags)
	assert.Equal(t, diag.FlatDiags{{
		Err:      "provider plugin unverified@v0.0.3 not registered at https://hub.cloudquery.io",
		Resource: "",
		Type:     diag.INTERNAL,
		Severity: diag.ERROR,
		Summary:  "failed to download providers: provider plugin unverified@v0.0.3 not registered at https://hub.cloudquery.io"}},
		diag.FlattenDiags(diags, true))

	_, diags = Download(context.Background(), pm, &DownloadOptions{
		Providers: []registry.Provider{
			{
				Name:    "unverified",
				Version: "v0.0.3",
				Source:  "cloudquery",
			},
		},
		NoVerify: true,
	})
	assert.Nil(t, diags)
}

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
