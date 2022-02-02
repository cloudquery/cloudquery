package registry

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/google/go-github/v35/github"
)

func TestCheckProviderUpdate(t *testing.T) {
	type githubResult struct {
		version string
		err     error
	}
	tests := []struct {
		name              string
		requestedProvider *config.RequiredProvider
		github            githubResult
		want              string
		wantErr           bool
	}{
		{
			"bad provider name",
			&config.RequiredProvider{Name: "very/strange/name"},
			githubResult{},
			"",
			true,
		},
		{
			"bad local provider version",
			&config.RequiredProvider{Name: "test", Version: "bad"},
			githubResult{},
			"",
			true,
		},
		{
			"github returns an error",
			&config.RequiredProvider{Name: "test", Version: "1.0.0"},
			githubResult{"", errors.New("fake")},
			"",
			true,
		},
		{
			"bad remote provider version",
			&config.RequiredProvider{Name: "test", Version: "1.0.0"},
			githubResult{"bad", nil},
			"",
			true,
		},
		{
			"remote version is newer",
			&config.RequiredProvider{Name: "test", Version: "1.0.0"},
			githubResult{"1.0.1", nil},
			"1.0.1",
			false,
		},
		{
			"versions are equal",
			&config.RequiredProvider{Name: "test", Version: "1.0.0"},
			githubResult{"1.0.0", nil},
			"",
			false,
		},
		{
			"local version is newer",
			&config.RequiredProvider{Name: "test", Version: "1.0.1"},
			githubResult{"1.0.0", nil},
			"",
			false,
		},
		{
			"latest version",
			&config.RequiredProvider{Name: "test", Version: "latest"},
			githubResult{"1.0.0", nil},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewRegistryHub("", WithLatestReleaseGetter(func(ctx context.Context, owner, repo string) (*github.RepositoryRelease, error) {
				return &github.RepositoryRelease{TagName: &tt.github.version}, tt.github.err
			}))
			got, err := h.CheckProviderUpdate(context.Background(), tt.requestedProvider)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hub.CheckProviderUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hub.CheckProviderUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
