package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	zerolog "github.com/rs/zerolog/log"
)

const (
	CloudQueryRegistryURL = "https://firestore.googleapis.com/v1/projects/hub-cloudquery/databases/(default)/documents/orgs/%s/providers/%s"

	// Timeout for http requests related to CloudQuery providers version check.
	versionCheckHTTPTimeout = time.Second * 10
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

	if organization != DefaultOrganization {
		if h.ProgressUpdater != nil {
			h.ProgressUpdater.Update(providerName, ui.StatusWarn, "skipped community provider verification...", 2)
		}
		return true
	}

	l := h.Logger.With("provider", providerName, "version", version)
	checksumsPath := filepath.Join(h.PluginDirectory, organization, providerName, version+".checksums.txt")
	checksumsURL := fmt.Sprintf("https://github.com/%s/%s/releases/latest/download/checksums.txt", organization, ProviderRepoName(providerName))
	if version != "latest" {
		checksumsURL = fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/checksums.txt", organization, ProviderRepoName(providerName), version)
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

// CheckProviderUpdate checks if there is an update available for the requested provider.
// Returns a new version if there is one, otherwise empty string.
// Call will be cancelled either if ctx is cancelled or after a timeout set by versionCheckHTTPTimeout.
// This function should not be called for a provider having Version set to "latest".
func (h Hub) CheckProviderUpdate(ctx context.Context, requestedProvider *config.RequiredProvider) (string, error) {
	organization, providerName, err := parseProviderSource(requestedProvider)
	if err != nil {
		return "", err
	}
	currentVersion, err := version.NewVersion(requestedProvider.Version)
	if err != nil {
		return "", fmt.Errorf("bad version: provider %s, version %s", providerName, requestedProvider.Version)
	}

	ctx, cancel := context.WithTimeout(ctx, versionCheckHTTPTimeout)
	defer cancel()
	latestVersion, err := h.getLatestRelease(ctx, organization, providerName)
	if err != nil {
		return "", err
	}
	v, err := version.NewVersion(latestVersion)
	if err != nil {
		return "", fmt.Errorf("bad version received: provider %s, version %s", providerName, latestVersion)
	}
	if currentVersion.LessThan(v) {
		return latestVersion, nil
	}
	return "", nil
}

func (h Hub) DownloadProvider(ctx context.Context, requestedProvider *config.RequiredProvider, noVerify bool) (ProviderDetails, error) {
	providerVersion := requestedProvider.Version
	organization, providerName, err := parseProviderSource(requestedProvider)
	if err != nil {
		return ProviderDetails{}, err
	}

	if providerVersion == "latest" {
		providerVersion, err = h.getLatestRelease(ctx, organization, providerName)
		if err != nil {
			return ProviderDetails{}, err
		}
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
		h.ProgressUpdater.Add(providerName, fmt.Sprintf("%s@%s", ProviderRepoName(providerName), providerVersion), providerVersion, 2)
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
		progressCB = ui.CreateProgressUpdater(h.ProgressUpdater, fmt.Sprintf("%s@%s", ProviderRepoName(providerName), providerVersion))
	}

	providerURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", organization, ProviderRepoName(providerName), providerVersion, getPluginBinaryName(providerName))
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

func (h Hub) getLatestRelease(ctx context.Context, organization, providerName string) (string, error) {
	versions, err := url.Parse(fmt.Sprintf(h.url+"/versions", organization, providerName))
	if err != nil {
		return "", err
	}
	qv := versions.Query()
	qv.Set("pageSize", "1")
	qv.Set("orderBy", "v_major desc, v_minor desc, v_patch desc, published_at desc")
	qv.Set("mask.fieldPaths", "tag")
	versions.RawQuery = qv.Encode()

	hc := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, versions.String(), nil)
	res, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var doc struct {
		Documents []struct {
			Name   string `json:"name"`
			Fields struct {
				Tag struct {
					Val string `json:"stringValue"`
				} `json:"tag"`
			} `json:"fields"`
		} `json:"documents"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return "", err
	}

	if len(doc.Documents) == 0 || doc.Documents[0].Fields.Tag.Val == "" {
		return "", fmt.Errorf("failed to find provider %s latest version", providerName)
	}
	return doc.Documents[0].Fields.Tag.Val, nil
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
	return fmt.Sprintf("%s_%s", ProviderRepoName(providerName), GetBinarySuffix())
}

func GetBinarySuffix() string {
	var suffix = ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return fmt.Sprintf("%s_%s%s", runtime.GOOS, runtime.GOARCH, suffix)
}

func parseProviderSource(requestedProvider *config.RequiredProvider) (string, string, error) {
	var requestedSource string
	if requestedProvider.Source == nil || *requestedProvider.Source == "" {
		requestedSource = requestedProvider.Name
	} else {
		requestedSource = *requestedProvider.Source
	}
	return ParseProviderName(requestedSource)
}
