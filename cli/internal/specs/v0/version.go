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
		usingVersion, latestVersion, outdated := warnIfOutdated(ctx, p, managedplugin.PluginSource, source.Metadata)
		if outdated && options.upgradeCheckOutput != nil {
			recommendUpgradeCheck(source.Name, usingVersion, latestVersion, options)
		}
	}
	for _, destination := range destinations {
		warnIfOutdated(ctx, p, managedplugin.PluginDestination, destination.Metadata)
	}
	for _, transformer := range transformers {
		warnIfOutdated(ctx, p, managedplugin.PluginTransformer, transformer.Metadata)
	}
}

func warnIfOutdated(ctx context.Context, p *managedplugin.PluginVersionWarner, kind managedplugin.PluginType, plugin Metadata) (usingVersion, latestVersion *semver.Version, outdated bool) {
	if plugin.Version == "" {
		return nil, nil, false
	}
	org, name, err := pluginPathToOrgName(plugin.Path)
	if err != nil {
		log.Debug().Str("plugin", plugin.Name).Err(err).Msg("failed to get org and name from plugin path")
		return nil, nil, false
	}
	// N.B.: warning is best-effort; we ignore errors, but the function still logs errors with Debug logs
	// We only check for outdated plugins if the registry is cloudquery or github and the org is cloudquery
	if plugin.Registry != RegistryCloudQuery && (plugin.Registry != RegistryGitHub || org != "cloudquery") {
		return nil, nil, false
	}
	usingVersion, err = semver.NewVersion(plugin.Version)
	if err != nil {
		log.Debug().Str("plugin", name).Str("version", plugin.Version).Err(err).Msg("failed to parse actual version")
		return nil, nil, false
	}
	latestVersion, err = p.LatestVersion(ctx, org, name, kind.String())
	if err != nil {
		return nil, nil, false
	}
	if !usingVersion.LessThan(latestVersion) {
		return nil, nil, false
	}
	log.Warn().
		Str("plugin", name).
		Str("using_version", usingVersion.String()).
		Str("latest_version", latestVersion.String()).
		Str("url", fmt.Sprintf("https://www.cloudquery.io/hub/plugins/%s/%s/%s", kind, org, name)).
		Msg("Plugin is outdated, consider upgrading to the latest version.")
	return usingVersion, latestVersion, true
}

func recommendUpgradeCheck(sourceName string, usingVersion, latestVersion *semver.Version, options versionWarningOptions) {
	majorsBehind := latestVersion.Major() - usingVersion.Major()
	if majorsBehind <= 0 {
		return
	}
	using := "v" + usingVersion.String()
	latest := "v" + latestVersion.String()
	command := upgradeCheckCommand(options.upgradeCheckConfigPaths, sourceName, latest)
	versionsWord := "versions"
	if majorsBehind == 1 {
		versionsWord = "version"
	}
	message := fmt.Sprintf("Source %s is %d major %s behind (%s → %s). Before you upgrade, run `%s` to see the schema impact on your destinations.", sourceName, majorsBehind, versionsWord, using, latest, command)
	log.Warn().
		Str("source", sourceName).
		Str("using_version", using).
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
