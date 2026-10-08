package specs

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
)

type versionWarningOptions struct {
	upgradeCheckConfigPaths []string
	upgradeCheckOutput      io.Writer
}

type VersionWarningOption func(*versionWarningOptions)

func WithUpgradeCheckRecommendation(configPaths []string, output io.Writer) VersionWarningOption {
	return func(o *versionWarningOptions) {
		o.upgradeCheckConfigPaths = configPaths
		o.upgradeCheckOutput = output
	}
}

func WarnOnOutdatedVersions(ctx context.Context, p *managedplugin.PluginVersionWarner, sources []*Source, destinations []*Destination, transformers []*Transformer, opts ...VersionWarningOption) {
	options := versionWarningOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	if p == nil {
		// This cannot happen at the time of writing. It could be nil if:
		// - The API base url is not set properly (but it's hardcoded)
		// - The options passed to getHubClient fail (but no options are passed)
		return // However, avoid panicking in case it's nil.
	}
	for _, source := range sources {
		org, name, err := pluginPathToOrgName(source.Path)
		if err != nil {
			log.Debug().Str("plugin", source.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		// We only check for outdated plugins if the registry is cloudquery or github and the org is cloudquery
		if source.Registry == RegistryCloudQuery || (source.Registry == RegistryGitHub && org == "cloudquery") {
			outdated, _ := p.WarnIfOutdated(ctx, org, name, managedplugin.PluginSource.String(), source.Version)
			if outdated && options.upgradeCheckOutput != nil {
				recommendUpgradeCheck(ctx, p, org, name, source, options)
			}
		}
	}
	for _, destination := range destinations {
		org, name, err := pluginPathToOrgName(destination.Path)
		if err != nil {
			log.Debug().Str("plugin", destination.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		// We only check for outdated plugins if the registry is cloudquery or github and the org is cloudquery
		if destination.Registry == RegistryCloudQuery || (destination.Registry == RegistryGitHub && org == "cloudquery") {
			_, _ = p.WarnIfOutdated(ctx, org, name, managedplugin.PluginDestination.String(), destination.Version)
		}
	}
	for _, transformer := range transformers {
		org, name, err := pluginPathToOrgName(transformer.Path)
		if err != nil {
			log.Debug().Str("plugin", transformer.Name).Err(err).Msg("failed to get org and name from plugin path")
			continue
		}
		// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
		// We only check for outdated plugins if the registry is cloudquery or github and the org is cloudquery
		if transformer.Registry == RegistryCloudQuery || (transformer.Registry == RegistryGitHub && org == "cloudquery") {
			_, _ = p.WarnIfOutdated(ctx, org, name, managedplugin.PluginTransformer.String(), transformer.Version)
		}
	}
}

func recommendUpgradeCheck(ctx context.Context, p *managedplugin.PluginVersionWarner, org, name string, source *Source, options versionWarningOptions) {
	latestVersion, err := p.LatestVersion(ctx, org, name, managedplugin.PluginSource.String())
	if err != nil {
		return
	}
	usingVersion, err := semver.NewVersion(source.Version)
	if err != nil {
		return
	}
	majorsBehind := latestVersion.Major() - usingVersion.Major()
	if majorsBehind <= 0 {
		return
	}
	latest := "v" + latestVersion.String()
	command := upgradeCheckCommand(options.upgradeCheckConfigPaths, source.Name, latest)
	versionsWord := "versions"
	if majorsBehind == 1 {
		versionsWord = "version"
	}
	message := fmt.Sprintf("Source %s is %d major %s behind (%s → %s). Before you upgrade, run `%s` to see the schema impact on your destinations.", source.Name, majorsBehind, versionsWord, source.Version, latest, command)
	log.Warn().
		Str("source", source.Name).
		Str("using_version", source.Version).
		Str("latest_version", latest).
		Str("command", command).
		Msg(message)
	_, _ = fmt.Fprintln(options.upgradeCheckOutput, message)
}

func upgradeCheckCommand(configPaths []string, sourceName, toVersion string) string {
	parts := make([]string, 0, len(configPaths)+7)
	parts = append(parts, "cloudquery", "upgrade", "check")
	for _, path := range configPaths {
		parts = append(parts, shellQuote(path))
	}
	parts = append(parts, "--source", shellQuote(sourceName), "--to", toVersion)
	return strings.Join(parts, " ")
}

func shellQuote(value string) string {
	if value == "" || strings.ContainsAny(value, " \t\n'\"") {
		return strconv.Quote(value)
	}
	return value
}

func pluginPathToOrgName(pluginPath string) (org string, name string, err error) {
	parts := strings.Split(pluginPath, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid plugin path: %s. format should be org/name", pluginPath)
	}
	return parts[0], parts[1], nil
}
