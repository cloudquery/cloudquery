package registry

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/stretchr/testify/assert"
)

func TestHub_CheckUpdate(t *testing.T) {
	testCases := []struct {
		Name          string
		Provider      Provider
		ExpectedError error
	}{
		{
			Name: "simple",
			Provider: Provider{
				Name:    "test",
				Version: "v0.0.5",
				Source:  DefaultOrganization,
			},
		},
		{
			Name: "bad-version",
			Provider: Provider{
				Name:    "test",
				Version: "va441.311.4123.444",
				Source:  DefaultOrganization,
			},
			ExpectedError: errors.New("bad version: test@va441.311.4123.444"),
		},
		{
			Name: "bad-provider-name",
			Provider: Provider{
				Name:    "bad-provider",
				Version: "v0.0.1",
				Source:  DefaultOrganization,
			},
			ExpectedError: errors.New("failed to find provider[bad-provider] latest version"),
		},
		{
			Name: "bad-org-name",
			Provider: Provider{
				Name:    "test",
				Version: "v0.0.1",
				Source:  "bad-org",
			},
			ExpectedError: errors.New("failed to find provider[test] latest version"),
		},
	}
	hub := NewRegistryHub(firebase.CloudQueryRegistryURL)

	latestVersion, err := hub.CheckUpdate(context.Background(), Provider{
		Name:    "test",
		Version: "v0.0.0",
		Source:  DefaultOrganization,
	})
	assert.Nil(t, err)
	assert.NotEqual(t, "", latestVersion)

	t.Run("check-update-with-latest", func(t *testing.T) {
		latestVersion, err := hub.CheckUpdate(context.Background(), Provider{
			Name:    "test",
			Version: latestVersion,
			Source:  DefaultOrganization,
		})
		assert.Nil(t, err)
		assert.Equal(t, "", latestVersion)
	})

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := hub.CheckUpdate(context.Background(), tc.Provider)
			if tc.ExpectedError != nil {
				assert.Equal(t, tc.ExpectedError, err)
				assert.Equal(t, "", result)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, latestVersion, result)
			}
		})
	}
}

func TestHub_Get(t *testing.T) {

	testCases := []struct {
		Name             string
		Provider         Provider
		ExpectedError    error
		ExpectedProvider ProviderBinary
	}{
		{
			Name: "missing_provider",
			Provider: Provider{
				Name:    "test2",
				Version: LatestVersion,
				Source:  "cloudqueryx",
			},
			ExpectedError:    errors.New("provider test2@v0.0.0 is missing, download it first"),
			ExpectedProvider: ProviderBinary{},
		},
		{
			Name: "existing_provider",
			Provider: Provider{
				Name:    "test",
				Version: LatestVersion,
				Source:  "hub",
			},
			ExpectedProvider: ProviderBinary{
				Provider: Provider{
					Name:    "test",
					Version: "v0.0.4",
					Source:  "cloudqueryx",
				},
				FilePath: "test/providers/cloudqueryx/test/v0.0.4",
			},
		},
		{
			Name: "existing_provider_w_version",
			Provider: Provider{
				Name:    "test",
				Version: "v0.0.1",
				Source:  "hub",
			},
			ExpectedProvider: ProviderBinary{
				Provider: Provider{
					Name:    "test",
					Version: "v0.0.1",
					Source:  "cloudqueryx",
				},
				FilePath: "test/providers/cloudqueryx/test/v0.0.1",
			},
		},
	}

	hub := NewRegistryHub(firebase.CloudQueryRegistryURL, WithPluginDirectory("test/providers"))
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := hub.Get(tc.Provider.Name, tc.Provider.Version)
			assert.Equal(t, tc.ExpectedError, err)
			assert.Equal(t, tc.ExpectedProvider.Provider, result.Provider)
			assert.Equal(t, tc.ExpectedProvider.FilePath, filepath.ToSlash(result.FilePath))
		})

	}
}

func TestHub_Download(t *testing.T) {
	t.Run("download-non-existing", func(t *testing.T) {
		hub := NewRegistryHub(firebase.CloudQueryRegistryURL)
		_, err := hub.Download(context.Background(), Provider{}, false)
		assert.Error(t, err)
	})
	t.Run("download-bad-version", func(t *testing.T) {
		hub := NewRegistryHub(firebase.CloudQueryRegistryURL, WithPluginDirectory(t.TempDir()))
		_, err := hub.Download(context.Background(), Provider{
			Name:    "test",
			Version: "v9.9.9",
			Source:  "cloudquery",
		}, false)
		assert.Error(t, err)
		_, err = hub.Get("test", "v0.0.11")
		assert.Error(t, err)

	})

	t.Run("download-test-provider", func(t *testing.T) {
		tempDir := t.TempDir()
		hub := NewRegistryHub(firebase.CloudQueryRegistryURL, WithPluginDirectory(tempDir))

		_, err := hub.Get("test", "v0.0.11")
		assert.Error(t, err)

		p, err := hub.Download(context.Background(), Provider{
			Name:    "test",
			Version: "v0.0.11",
			Source:  DefaultOrganization,
		}, false)
		assert.Nil(t, err)
		assert.True(t, strings.HasPrefix(p.FilePath, tempDir))
		assert.Equal(t, Provider{
			Name:    "test",
			Version: "v0.0.11",
			Source:  DefaultOrganization,
		}, p.Provider)

		pGet, err := hub.Get("test", "v0.0.11")
		assert.Nil(t, err)
		assert.Equal(t, p, pGet)
	})

}
