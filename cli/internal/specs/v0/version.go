package specs

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
)

func WarnOnOutdatedVersions(ctx context.Context, p *managedplugin.PluginVersionWarner, sources []*Source, destinations []*Destination, transformers []*Transformer) {
	for _, source := range sources {
		org, name, err := pluginPathToOrgName(source.Path)
		if err != nil {
			log.Debug().Str("plugin", source.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		_, _ = p.WarnIfOutdated(ctx, org, name, managedplugin.PluginSource.String(), source.Version)
	}
	for _, destination := range destinations {
		org, name, err := pluginPathToOrgName(destination.Path)
		if err != nil {
			log.Debug().Str("plugin", destination.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		_, _ = p.WarnIfOutdated(ctx, org, name, managedplugin.PluginDestination.String(), destination.Version)
	}
	for _, transformer := range transformers {
		org, name, err := pluginPathToOrgName(transformer.Path)
		if err != nil {
			log.Debug().Str("plugin", transformer.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		_, _ = p.WarnIfOutdated(ctx, org, name, managedplugin.PluginTransformer.String(), transformer.Version)
	}
}

func pluginPathToOrgName(pluginPath string) (org string, name string, err error) {
	parts := strings.Split(pluginPath, "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid plugin path: %s", pluginPath)
	}
	return parts[0], parts[1], nil
}
