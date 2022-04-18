package core

import (
	"context"
	"time"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

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
func Download(ctx context.Context, manager *plugin.Manager, opts *DownloadOptions) (*DownloadResult, diag.Diagnostics) {
	log.Info().Interface("providers", opts.Providers).Msg("downloading providers")
	startTime := time.Now()
	downloaded, err := manager.DownloadProviders(ctx, opts.Providers, opts.NoVerify)
	if err != nil {
		return nil, diag.Diagnostics{diag.NewBaseError(err, diag.INTERNAL, diag.WithSeverity(diag.ERROR), diag.WithSummary("failed to download providers"))}
	}
	log.Info().Interface("providers", opts.Providers).Dur("duration", time.Since(startTime)).Msg("providers download successfully")
	return &DownloadResult{downloaded}, nil
}
