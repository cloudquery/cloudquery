package registry

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/internal/versions"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// Timeout for http requests related to CloudQuery providers version check.
	versionCheckHTTPTimeout = time.Second * 10
)

type Hub struct {
	// Optional: Where to save downloaded providers, by default current working directory, defaults to ./cq/providers
	PluginDirectory string
	// Optional: Download propagator allows the creator to get called back on download progress and completion.
	ProgressUpdater ui.Progress
	// Url for hub to connect to download and verify plugins
	url string
	// map of downloaded providers
	providers map[string]ProviderBinary
}

type Option func(h *Hub)

func WithPluginDirectory(d string) Option {
	return func(h *Hub) {
		h.PluginDirectory = d
	}
}

func WithProgress(u ui.Progress) Option {
	return func(h *Hub) {
		h.ProgressUpdater = u
	}
}

func NewRegistryHub(url string, opts ...Option) *Hub {
	h := &Hub{
		PluginDirectory: filepath.Join(".", ".cq", "providers"),
		url:             url,
		providers:       make(map[string]ProviderBinary),
	}
	// apply the list of options to hub
	for _, opt := range opts {
		opt(h)
	}
	h.loadExisting()
	return h
}

// Get returns a loaded provider from the hub without downloading it again, returns an error if not found.
func (h Hub) Get(providerName, providerVersion string) (ProviderBinary, error) {
	if providerVersion == "latest" {
		latestVersion, _ := version.NewVersion("v0.0.0")
		for _, p := range h.providers {
			if p.Name != providerName {
				continue
			}
			currentVersion, err := version.NewVersion(p.Version)
			if err != nil {
				log.Warn().Str("provider", providerName).Str("version", providerVersion).Msg("bad version provider exists in directory")
				continue
			}
			if currentVersion.GreaterThan(latestVersion) {
				latestVersion = currentVersion
			}
		}
		providerVersion = latestVersion.Original()
	}
	// TODO: support organization naming level for providers
	pd, ok := h.providers[fmt.Sprintf("%s-%s", providerName, providerVersion)]
	if !ok {
		return ProviderBinary{}, fmt.Errorf("provider %s@%s is missing, download it first", providerName, providerVersion)
	}
	return pd, nil
}

// CheckUpdate checks if there is an update available for the requested provider.
// Returns a new version if there is one, otherwise empty string.
// Call will be cancelled either if ctx is cancelled or after a timeout set by versionCheckHTTPTimeout.
// This function should not be called for a provider having Version set to "latest".
func (Hub) CheckUpdate(ctx context.Context, provider Provider) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, versionCheckHTTPTimeout)
	defer cancel()
	latestVersion, err := getLatestRelease(ctx, provider.Source, provider.Name)
	if err != nil {
		return "", err
	}
	if provider.Version == LatestVersion {
		return latestVersion, nil
	}
	v, err := version.NewVersion(latestVersion)
	if err != nil {
		return "", fmt.Errorf("bad version received: provider %s, version %s", provider.Name, latestVersion)
	}
	currentVersion, err := version.NewVersion(provider.Version)
	if err != nil {
		return "", fmt.Errorf("bad version: %s", provider)
	}
	if currentVersion.LessThan(v) {
		return latestVersion, nil
	}
	return "", nil
}

func (h Hub) Download(ctx context.Context, provider Provider, noVerify bool) (ProviderBinary, error) {
	var (
		requestedVersion = provider.Version
		err              error
	)
	if requestedVersion == "latest" {
		requestedVersion, err = getLatestRelease(ctx, provider.Source, provider.Name)
		if err != nil {
			return ProviderBinary{}, err
		}
	}
	p, ok := h.providers[fmt.Sprintf("%s-%s", provider.Name, requestedVersion)]
	if !ok {
		return h.downloadProvider(ctx, provider, requestedVersion, noVerify)
	}
	if p.Version != requestedVersion {
		log.Info().Str("current", p.Version).Str("requested", requestedVersion).Msg("current version is not as requested version updating provider")
		return h.downloadProvider(ctx, provider, requestedVersion, noVerify)
	}

	if h.ProgressUpdater != nil {
		// set up a done download progress
		h.ProgressUpdater.Add(provider.Name, fmt.Sprintf("%s@%s", ProviderRepoName(provider.Name), requestedVersion), requestedVersion, 2)
	}

	if noVerify {
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(provider.Name, ui.StatusWarn, "skipped verification...", 2)
		}
		return p, nil
	}

	if !h.verifyProvider(ctx, provider, requestedVersion) {
		return ProviderBinary{}, fmt.Errorf("provider %s@%s verification failed", provider.Name, requestedVersion)
	}
	return p, nil
}

func (h Hub) verifyProvider(ctx context.Context, provider Provider, version string) bool {
	if provider.Source != DefaultOrganization {
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(provider.Name, ui.StatusWarn, "skipped community provider verification...", 2)
		}
		return true
	}

	l := log.With().Str("provider", provider.Name).Str("version", version).Logger()
	if h.ProgressUpdater != nil {
		h.ProgressUpdater.Update(provider.Name, ui.StatusInProgress, "Verifying...", 1)
	}

	checksumsPath, err := h.downloadFile(ctx, l, provider, version, "checksums.txt")
	if err != nil {
		l.Error().Err(err).Msg("failed to download checksums file")
		return false
	}

	_, err = h.downloadFile(ctx, l, provider, version, "checksums.txt.sig")
	if err != nil {
		l.Error().Err(err).Msg("failed to download signature file")
		return false
	}

	err = validateFile(checksumsPath, checksumsPath+".sig")
	if err != nil {
		l.Error().Err(err).Msg("validating provider signature failed")
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(provider.Name, ui.StatusError, "Bad signature", 0)
		}
		return false
	}
	providerPath := h.getProviderPath(provider.Source, provider.Name, version)
	if err = validateChecksumProvider(providerPath, checksumsPath); err != nil {
		l.Error().Err(err).Msg("validating provider checksum failed")
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(provider.Name, ui.StatusError, "Bad checksum", 0)
		}
		return false
	}
	if h.ProgressUpdater != nil {
		h.ProgressUpdater.Update(provider.Name, ui.StatusOK, "verified", 1)
	}
	return true
}

func (h Hub) downloadFile(ctx context.Context, l zerolog.Logger, provider Provider, version string, fileName string) (string, error) {
	path := filepath.Join(h.PluginDirectory, provider.Source, provider.Name, version+"."+fileName)
	osFs := file.NewOsFs()
	switch provider.Source {
	case DefaultOrganization:
		// handle CloudQuery monorepo checksums, which by necessity uses a different
		// convention from community plugins. This version of the CLI only supports "source" type plugins,
		// but support for more will be added in the future.
		tag := fmt.Sprintf("plugins/source/%s/%s", provider.Name, version)
		checksumsURL := fmt.Sprintf("https://github.com/%s/cloudquery/releases/download/%s/%s", DefaultOrganization, tag, fileName)
		l.Debug().Str("url", checksumsURL).Str("path", path).Msg("downloading checksums file from monorepo")
		err := osFs.DownloadFile(ctx, path, checksumsURL, nil)
		if err == nil {
			return path, nil
		}
		// if we failed to download checksums from the monorepo, it might be because the plugin version
		// was released on the older provider repo. Fall through to try there before giving up
		fallthrough
	default:
		checksumsURL := fmt.Sprintf("https://github.com/%s/%s/releases/latest/download/%s", provider.Source, ProviderRepoName(provider.Name), fileName)
		if version != "latest" {
			checksumsURL = fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", provider.Source, ProviderRepoName(provider.Name), version, fileName)
		}
		l.Debug().Str("url", checksumsURL).Str("path", path).Msg("downloading checksums file")
		err := osFs.DownloadFile(ctx, path, checksumsURL, nil)
		return path, err
	}
}

func (h Hub) downloadProvider(ctx context.Context, provider Provider, requestedVersion string, noVerify bool) (ProviderBinary, error) {
	if !h.verifyRegistered(provider.Source, provider.Name, requestedVersion, noVerify) {
		return ProviderBinary{}, fmt.Errorf("provider plugin %s@%s not registered at https://hub.cloudquery.io", provider.Name, requestedVersion)
	}
	// build fully qualified plugin directory for given plugin
	pluginDir := filepath.Join(h.PluginDirectory, provider.Source, provider.Name)
	osFs := file.NewOsFs()
	if err := osFs.MkdirAll(pluginDir, os.ModePerm); err != nil {
		return ProviderBinary{}, diag.FromError(err, diag.USER, diag.WithSummary("failed to create plugin directory"))
	}
	// Create a new progress updater callback func
	var progressCB ui.ProgressUpdateFunc
	if h.ProgressUpdater != nil {
		progressCB = ui.CreateProgressUpdater(h.ProgressUpdater, fmt.Sprintf("%s@%s", ProviderRepoName(provider.Name), requestedVersion))
	}

	var err error
	providerPath := h.getProviderPath(provider.Source, provider.Name, requestedVersion)

	switch provider.Source {
	case DefaultOrganization:
		// we use a special convention for the CloudQuery monorepo
		providerURL := fmt.Sprintf("https://github.com/%s/cloudquery/releases/download/%s/%s", DefaultOrganization, getMonorepoPluginTag("source", provider.Name, requestedVersion), getPluginBinaryName(provider.Name))
		err = osFs.DownloadFile(ctx, providerPath, providerURL, progressCB)
		if err == nil {
			break
		}
		// if the download was attempted from the monorepo but failed, it could be that
		// the specified version isn't available there (but is on the original provider repo).
		// In this case, we retry the download from the original provider repo before reporting
		// a failure.
		fallthrough
	default:
		providerURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", provider.Source, ProviderRepoName(provider.Name), requestedVersion, getPluginBinaryName(provider.Name))
		err = osFs.DownloadFile(ctx, providerPath, providerURL, progressCB)
	}
	if err != nil {
		return ProviderBinary{}, fmt.Errorf("plugin %s/%s@%s failed to download: %w", provider.Source, provider.Name, requestedVersion, err)
	}
	
	//if ok := h.verifyProvider(ctx, provider, requestedVersion); !ok {
	//	return ProviderBinary{}, fmt.Errorf("plugin %s/%s@%s failed to verify", provider.Source, provider.Name, requestedVersion)
	//}

	if err := osFs.Chmod(providerPath, 0754); err != nil {
		return ProviderBinary{}, err
	}

	details := ProviderBinary{
		Provider: Provider{
			Name:    provider.Name,
			Version: requestedVersion,
			Source:  provider.Source,
		},
		FilePath: providerPath,
	}
	h.providers[fmt.Sprintf("%s-%s", provider.Name, requestedVersion)] = details

	return details, nil
}

func getLatestRelease(ctx context.Context, organization, providerName string) (string, error) {
	// Only "source" type plugins are supported in this version of the CLI. This will be
	// expanded to other types in the future.
	v, err := versions.NewClient().GetLatestProviderRelease(ctx, organization, "source", providerName)
	if err != nil {
		return "", fmt.Errorf("failed to find provider[%s] latest version", providerName)
	}
	return v, nil
}

func (h Hub) verifyRegistered(organization, providerName, version string, noVerify bool) bool {
	if noVerify {
		log.Warn().Str("provider", providerName).Msg("skipping plugin registry verification")
		return true
	}
	log.Debug().Str("provider", providerName).Str("version", version).Msg("verifying provider plugin is registered")
	if !h.isProviderRegistered(organization, providerName) {
		return false
	}

	log.Debug().Str("provider", providerName).Str("version", version).Msg("provider plugin is registered")
	return true
}

func (Hub) isProviderRegistered(org, provider string) bool {
	client := firebase.New(firebase.CloudQueryRegistryURL)
	return client.IsProviderRegistered(org, provider)
}

// GetProviderPath returns expected path of provider on file system from name and version of plugin
func (h Hub) getProviderPath(org, name, version string) string {
	return filepath.Join(h.PluginDirectory, org, name, fmt.Sprintf("%s-%s", version, GetBinarySuffix()))
}

func (h Hub) loadExisting() {
	osFs := file.NewOsFs()
	_ = osFs.WalkPathTree(h.PluginDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Debug().Str("directory", h.PluginDirectory).Msg("failed to read plugin directory, no existing plugins loaded")
			return nil
		}
		if info.IsDir() {
			return nil
		}
		// skip checksum files, they will be downloaded again
		if strings.Contains(info.Name(), "checksums") {
			return nil
		}
		provider := filepath.Base(filepath.Dir(path))
		if strings.HasSuffix(path, ".tmp") {
			log.Debug().Str("provider", provider).Msg("found temp provider file, cleaning up")
			if err := osFs.Remove(path); err != nil {
				log.Warn().Str("provider", provider).Msg("failed to remove temp provider file")
			}
			return nil
		}
		organization := filepath.Base(filepath.Dir(filepath.Dir(path)))
		pVersion := strings.Split(filepath.Base(path), "-"+GetBinarySuffix())[0]

		h.providers[fmt.Sprintf("%s-%s", provider, pVersion)] = ProviderBinary{
			Provider: Provider{
				Name:    provider,
				Version: pVersion,
				Source:  organization,
			},
			FilePath: path,
		}
		log.Debug().Str("provider", provider).Str("version", pVersion).Msg("found existing provider")
		return nil
	})
}

// getPluginBinaryName returns fully qualified CloudQuery plugin name based on running OS
func getPluginBinaryName(providerName string) string {
	return fmt.Sprintf("%s_%s", ProviderRepoName(providerName), GetBinarySuffix())
}

// getMonorepoPluginTag returns
func getMonorepoPluginTag(pluginType, pluginName, version string) string {
	return fmt.Sprintf("plugins/%s/%s/%s", pluginType, pluginName, version)
}

func GetBinarySuffix() string {
	var suffix = ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return fmt.Sprintf("%s_%s%s", runtime.GOOS, runtime.GOARCH, suffix)
}
