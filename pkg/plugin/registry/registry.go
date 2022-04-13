package registry

import (
	"context"
	"fmt"
)

const LatestVersion = "latest"

type ProviderBinary struct {
	Provider
	FilePath string
}

type Provider struct {
	// Name of the provider
	Name string
	// Version of the provider we want to download
	Version string
	// Source from where we want to download the provider from
	Source string
}

func (p Provider) String() string {
	return fmt.Sprintf("%s@%s", p.Name, p.Version)
}

//go:generate mockgen -package=registry -destination=./mock_registry.go . Registry
type Registry interface {
	// Get returns a loaded provider from the hub without downloading it again, returns an error if not found
	Get(providerName, providerVersion string) (ProviderBinary, error)
	// CheckUpdate checks if there is an update available for the requested provider.
	CheckUpdate(ctx context.Context, provider Provider) (string, error)
	// Download downloads a single provider from remote source.
	Download(ctx context.Context, provider Provider, noVerify bool) (ProviderBinary, error)
}
