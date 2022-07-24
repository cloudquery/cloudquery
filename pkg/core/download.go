package core

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/rs/zerolog/log"
)

type DownloadOptions struct {
	// Providers to purge data from, the provider name should be the plugin name
	Providers []registry.Provider
	// Whether download should verify plugins after they are downloaded
	NoVerify bool
}

// DownloadResult output from Download command
type DownloadResult struct {
	// Downloaded is a list of downloaded providers
	Downloaded []registry.ProviderBinary
}

// Download one or more providers from remote registry
func Download(ctx context.Context, manager *plugin.Manager, opts *DownloadOptions) (*DownloadResult, error) {
	log.Info().Interface("providers", opts.Providers).Msg("downloading providers")
	startTime := time.Now()
	downloaded, err := manager.DownloadProviders(ctx, opts.Providers, opts.NoVerify)
	if err != nil {
		return nil, fmt.Errorf("failed to download providers: %w", err)
	}
	log.Info().Interface("providers", opts.Providers).Dur("duration", time.Since(startTime)).Msg("providers download successfully")
	return &DownloadResult{downloaded}, nil
}
