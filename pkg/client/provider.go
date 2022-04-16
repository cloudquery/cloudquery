package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/rs/zerolog/log"
)

type GetProviderConfigOptions struct {
	Provider registry.Provider
}

func GetProviderConfiguration(ctx context.Context, pm *plugin.Manager, opts *GetProviderConfigOptions) (*cqproto.GetProviderConfigResponse, diag.Diagnostics) {
	providerPlugin, err := pm.CreatePlugin(&plugin.CreationOptions{Provider: opts.Provider})
	if err != nil {
		log.Error().Err(err).Str("provider", opts.Provider.Name).Str("version", opts.Provider.Version).Msg("failed to create provider plugin")
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	defer pm.ClosePlugin(providerPlugin)
	result, err := providerPlugin.Provider().GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
	if err != nil {
		return result, diag.FromError(err, diag.INTERNAL)
	}
	return result, nil
}

type TestOptions struct {
	Connection   cqproto.ConnectionDetails
	Config       []byte
	CreationInfo *plugin.CreationOptions
}

// TODO: add tests for Test method, add a "special" configuration that will return a failure in test provider

// Test checks if a provider's configure will work, this method is usually used to check that the credentials / provider configuration
// is correct works.
func Test(ctx context.Context, pm *plugin.Manager, opts TestOptions) (bool, error) {
	p, err := pm.CreatePlugin(opts.CreationInfo)
	if err != nil {
		return false, err
	}
	defer pm.ClosePlugin(p)
	log.Info().Str("provider", opts.CreationInfo.Provider.Name).Str("version", opts.CreationInfo.Provider.Version).Msg("requesting provider to configure")
	// TODO: check configure provider response errors/diagnostics
	_, err = p.Provider().ConfigureProvider(ctx, &cqproto.ConfigureProviderRequest{
		CloudQueryVersion: Version,
		Connection:        opts.Connection,
		Config:            opts.Config,
	})
	if err != nil {
		return false, fmt.Errorf("provider test connection failed. Reason: %w", err)
	}
	return true, nil
}

type CheckUpdatesOptions struct {
	Providers []registry.Provider
}

type AvailableUpdate struct {
	Name             string
	CurrentVersion   string
	AvailableVersion string
}

// TODO: support installed providers table to save what providers were installed regardless of current disk

// CheckAvailableUpdates checks if any updates are available for providers, if a provider's version is set to latest,
// update will check vs "latest" available provider located in the local disk.
func CheckAvailableUpdates(ctx context.Context, reg registry.Registry, opts *CheckUpdatesOptions) ([]AvailableUpdate, diag.Diagnostics) {
	var (
		diags   diag.Diagnostics
		updates = make([]AvailableUpdate, 0, len(opts.Providers))
	)
	for _, p := range opts.Providers {
		var version = p.Version
		if p.Version == registry.LatestVersion {
			pb, err := reg.Get(p.Name, p.Version)
			// This can happen, when we check for updates, but we don't have any providers downloaded, in this case
			// the latest provider will should be downloaded via the Download command.
			if err != nil {
				continue
			}
			version = pb.Version
		}
		log.Info().Str("provider", p.Name).Str("version", version).Msg("checking update for provider")
		updateVersion, err := reg.CheckUpdate(ctx, p)
		if err != nil {
			log.Error().Err(err).Str("provider", p.Name).Str("version", version).Msg("failed to check provider update")
			diags = diags.Add(diag.FromError(err, diag.INTERNAL))
		}
		if updateVersion == "" {
			continue
		}
		log.Info().Str("provider", p.Name).Str("version", p.Version).Str("new_version", updateVersion).Msg("update available for provider")
		updates = append(updates, AvailableUpdate{
			Name:             p.Name,
			CurrentVersion:   version,
			AvailableVersion: updateVersion,
		})
	}
	return updates, diags
}
