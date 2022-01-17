package registry

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	zerolog "github.com/rs/zerolog/log"
)

const (
	CloudQueryRegistryURl = "https://firestore.googleapis.com/v1/projects/hub-cloudquery/databases/(default)/documents/orgs/%s/providers/%s"
)

type ProviderDetails struct {
	Name         string
	Version      string
	Organization string
	FilePath     string
}

type Hub struct {
	// Optional: Where to save downloaded providers, by default current working directory, defaults to ./cq/providers
	PluginDirectory string
	// Optional: Download propagator allows the creator to get called back on download progress and completion.
	ProgressUpdater ui.Progress
	// Optional: logger to use, if not defined global logger is used.
	Logger hclog.Logger
	// Url for hub to connect to download and verify plugins
	url string
	// map of downloaded providers
	providers map[string]ProviderDetails
}

const (
	providerDisplayMsg = "cq-provider-%s@%s"
)

type Option func(h *Hub)

func NewRegistryHub(url string, opts ...Option) *Hub {
	h := &Hub{
		PluginDirectory: filepath.Join(".", ".cq", "providers"),
		Logger:          logging.NewZHcLog(&zerolog.Logger, ""),
		url:             url,
		providers:       make(map[string]ProviderDetails),
	}
	// apply the list of options to hub
	for _, opt := range opts {
		opt(h)
	}
	h.loadExisting()
	return h
}

// GetProvider returns a loaded provider from the hub without downloading it again, returns an error if not found.
func (h Hub) GetProvider(providerName, providerVersion string) (ProviderDetails, error) {
	if providerVersion == "latest" {
		latestVersion, _ := version.NewVersion("v0.0.0")
		for _, p := range h.providers {
			if p.Name != providerName {
				continue
			}
			currentVersion, err := version.NewVersion(p.Version)
			if err != nil {
				h.Logger.Warn("bad version provider exists in directory", "provider", p.Name, "version", p.Version)
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
		return ProviderDetails{}, fmt.Errorf("provider %s@%s is missing, download it first", providerName, providerVersion)
	}
	return pd, nil
}

func (h Hub) VerifyProvider(ctx context.Context, organization, providerName, version string) bool {

	if organization != defaultOrganization {
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(providerName, ui.StatusWarn, "skipped community provider verification...", 2)
		}
		return true
	}

	l := h.Logger.With("provider", providerName, "version", version)
	checksumsPath := filepath.Join(h.PluginDirectory, organization, providerName, version+".checksums.txt")
	checksumsURL := fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/latest/download/checksums.txt", organization, providerName)
	if version != "latest" {
		checksumsURL = fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/download/%s/checksums.txt", organization, providerName, version)
	}
	if h.ProgressUpdater != nil {
		h.ProgressUpdater.Update(providerName, ui.StatusInProgress, "Verifying...", 1)
	}
	l.Debug("downloading checksums file", "url", checksumsURL, "path", checksumsPath)
	// download checksums
	osFs := file.NewOsFs()
	if err := osFs.DownloadFile(ctx, checksumsPath, checksumsURL, nil); err != nil {
		l.Error("failed to download checksums file", "providerName", providerName, "error", err)
		return false
	}
	l.Debug("downloading checksums signature", "url", checksumsURL, "path", checksumsPath)
	// download checksums signature
	if err := osFs.DownloadFile(ctx, checksumsPath+".sig", checksumsURL+".sig", nil); err != nil {
		l.Error("failed to download signature file", "providerName", providerName, "error", err)
		return false
	}
	err := validateFile(checksumsPath, checksumsPath+".sig")
	if err != nil {
		l.Error("validating provider signature failed", "providerName", providerName, "error", err)
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(providerName, ui.StatusError, "Bad signature", 0)
		}
		return false
	}
	providerPath := h.getProviderPath(organization, providerName, version)
	if err = validateChecksumProvider(providerPath, checksumsPath); err != nil {
		l.Error("validating provider checksum failed", "providerName", providerName, "error", err)
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(providerName, ui.StatusError, "Bad checksum", 0)
		}
		return false
	}
	if h.ProgressUpdater != nil {
		h.ProgressUpdater.Update(providerName, ui.StatusOK, "verified", 1)
	}
	return true
}

// CheckProviderUpdate - checks if there is an update for provider, returns nil if there is no update
func (h Hub) CheckProviderUpdate(ctx context.Context, requestedProvider *config.RequiredProvider) (*string, error) {
	organization, providerName, err := ParseProviderName(requestedProvider.Name)
	if err != nil {
		return nil, err
	}
	currentVersion, err := version.NewVersion(requestedProvider.Version)
	if err != nil {
		return nil, fmt.Errorf("bad version: provider %s, version %s", providerName, requestedProvider.Version)
	}

	release, err := h.getRelease(ctx, organization, providerName, "latest")
	if err != nil {
		return nil, err
	}
	latestVersion := release.GetTagName()
	v, err := version.NewVersion(latestVersion)
	if err != nil {
		return nil, fmt.Errorf("bad version received: provider %s, version %s", providerName, latestVersion)
	}
	if currentVersion.LessThan(v) {
		return &latestVersion, nil
	}
	return nil, nil
}

func (h Hub) DownloadProvider(ctx context.Context, requestedProvider *config.RequiredProvider, noVerify bool) (ProviderDetails, error) {

	providerVersion := requestedProvider.Version
	organization, providerName, err := ParseProviderName(requestedProvider.Name)
	if err != nil {
		return ProviderDetails{}, err
	}

	if providerVersion == "latest" {
		release, err := h.getRelease(ctx, organization, providerName, providerVersion)
		if err != nil {
			return ProviderDetails{}, err
		}
		providerVersion = release.GetTagName()
	}
	p, ok := h.providers[fmt.Sprintf("%s-%s", providerName, providerVersion)]
	if !ok {
		return h.downloadProvider(ctx, organization, providerName, providerVersion, noVerify)
	}
	if p.Version != providerVersion {
		h.Logger.Info("Current version is not as requested version updating provider", "current", p.Version, "requested", providerVersion)
		return h.downloadProvider(ctx, organization, providerName, providerVersion, noVerify)
	}

	if h.ProgressUpdater != nil {
		// Setup a done download progress
		h.ProgressUpdater.Add(providerName, fmt.Sprintf("cq-provider-%s@%s", providerName, providerVersion), providerVersion, 2)
	}

	if noVerify {
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(providerName, ui.StatusWarn, "skipped verification...", 2)
		}
		return p, nil
	}

	if !h.VerifyProvider(ctx, organization, providerName, providerVersion) {
		return ProviderDetails{}, fmt.Errorf("provider %s@%s verification failed", providerName, providerVersion)
	}
	return p, nil
}

func (h Hub) downloadProvider(ctx context.Context, organization, providerName, providerVersion string, noVerify bool) (ProviderDetails, error) {

	if !h.verifyRegistered(organization, providerName, providerVersion, noVerify) {
		return ProviderDetails{}, fmt.Errorf("provider plugin %s@%s not registered at https://hub.cloudquery.io", providerName, providerVersion)
	}
	// build fully qualified plugin directory for given plugin
	pluginDir := filepath.Join(h.PluginDirectory, organization, providerName)
	osFs := file.NewOsFs()
	if err := osFs.MkdirAll(pluginDir, os.ModePerm); err != nil {
		return ProviderDetails{}, err
	}

	// Create a new progress updater callback func
	var progressCB ui.ProgressUpdateFunc
	if h.ProgressUpdater != nil {
		progressCB = ui.CreateProgressUpdater(h.ProgressUpdater, fmt.Sprintf(providerDisplayMsg, providerName, providerVersion))
	}

	providerURL := fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/download/%s/%s", organization, providerName, providerVersion, getPluginBinaryName(providerName))
	providerPath := h.getProviderPath(organization, providerName, providerVersion)
	if err := osFs.DownloadFile(ctx, providerPath, providerURL, progressCB); err != nil {
		return ProviderDetails{}, fmt.Errorf("plugin %s/%s@%s failed to download: %w", organization, providerName, providerVersion, err)
	}

	if ok := h.VerifyProvider(ctx, organization, providerName, providerVersion); !ok {
		return ProviderDetails{}, fmt.Errorf("plugin %s/%s@%s failed to verify", organization, providerName, providerVersion)
	}

	if err := osFs.Chmod(providerPath, 0754); err != nil {
		return ProviderDetails{}, err
	}

	details := ProviderDetails{
		Name:         providerName,
		Version:      providerVersion,
		Organization: organization,
		FilePath:     providerPath,
	}
	h.providers[fmt.Sprintf("%s-%s", providerName, providerVersion)] = details

	return details, nil
}

func (h Hub) getRelease(ctx context.Context, organization, providerName, version string) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	if version != "latest" {
		release, _, err := client.Repositories.GetReleaseByTag(ctx, organization, fmt.Sprintf("cq-provider-%s", providerName), version)
		return release, err
	}
	release, _, err := client.Repositories.GetLatestRelease(ctx, organization, fmt.Sprintf("cq-provider-%s", providerName))
	return release, err
}

func (h Hub) getLatestReleaseVersion(ctx context.Context, organization, providerName string) (*version.Version, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, organization, fmt.Sprintf("cq-provider-%s", providerName))
	if err != nil {
		return nil, err
	}
	return version.NewVersion(release.GetTagName())
}

func (h Hub) verifyRegistered(organization, providerName, version string, noVerify bool) bool {
	if noVerify {
		h.Logger.Warn("skipping plugin registry verification", "provider", providerName)
		return true
	}
	h.Logger.Debug("verifying provider plugin is registered", "provider", providerName, "version", version)
	if !h.isProviderRegistered(organization, providerName) {
		return false
	}

	h.Logger.Debug("provider plugin is registered", "provider", providerName, "version", version)
	return true
}

func (h Hub) isProviderRegistered(org, provider string) bool {
	url := fmt.Sprintf(h.url, org, provider)
	res, err := http.Get(url)
	if err != nil {
		h.Logger.Error("failed to check if provider is registered", "error", err)
		return false
	}
	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return false
		default:
			return false
		}
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	return true
}

// GetProviderPath returns expected path of provider on file system from name and version of plugin
func (h Hub) getProviderPath(org, name, version string) string {
	return filepath.Join(h.PluginDirectory, org, name, fmt.Sprintf("%s-%s", version, GetBinarySuffix()))
}

func (h Hub) loadExisting() {
	osFs := file.NewOsFs()
	_ = osFs.WalkPathTree(h.PluginDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			h.Logger.Debug("failed to read plugin directory, no existing plugins loaded", "directory", h.PluginDirectory)
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
			h.Logger.Debug("found temp provider file, cleaning up", "provider", provider)
			if err := osFs.Remove(path); err != nil {
				h.Logger.Warn("failed to remove temp provider file", "provider", provider)
			}
			return nil
		}
		organization := filepath.Base(filepath.Dir(filepath.Dir(path)))
		pVersion := strings.Split(filepath.Base(path), "-"+GetBinarySuffix())[0]

		h.providers[fmt.Sprintf("%s-%s", provider, pVersion)] = ProviderDetails{
			Name:         provider,
			Version:      pVersion,
			Organization: organization,
			FilePath:     path,
		}
		h.Logger.Debug("found existing provider", "provider", provider, "version", pVersion)
		return nil
	})
}

// getPluginBinaryName returns fully qualified CloudQuery plugin name based on running OS
func getPluginBinaryName(providerName string) string {
	return fmt.Sprintf("cq-provider-%s_%s", providerName, GetBinarySuffix())
}

func GetBinarySuffix() string {
	var suffix = ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return fmt.Sprintf("%s_%s%s", runtime.GOOS, runtime.GOARCH, suffix)
}
