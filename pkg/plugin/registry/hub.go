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

	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/hashicorp/go-version"
)

const (
	// CloudQueryRegistryURL default CloudQuery registry Hub URL
	CloudQueryRegistryURL = "https://firestore.googleapis.com/v1/projects/hub-cloudquery/databases/(default)/documents/orgs/%s/providers/%s"

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
func (h Hub) CheckUpdate(ctx context.Context, provider Provider) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, versionCheckHTTPTimeout)
	defer cancel()
	latestVersion, err := h.getLatestRelease(ctx, provider.Source, provider.Name)
	if err != nil {
		return "", err
	}
	v, err := version.NewVersion(latestVersion)
	if err != nil {
		return "", fmt.Errorf("bad version received: provider %s, version %s", provider.Name, latestVersion)
	}
	if provider.Version == LatestVersion {
		return latestVersion, nil
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
		requestedVersion, err = h.getLatestRelease(ctx, provider.Source, provider.Name)
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
	checksumsPath := filepath.Join(h.PluginDirectory, provider.Source, provider.Name, version+".checksums.txt")
	checksumsURL := fmt.Sprintf("https://github.com/%s/%s/releases/latest/download/checksums.txt", provider.Source, ProviderRepoName(provider.Name))
	if version != "latest" {
		checksumsURL = fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/checksums.txt", provider.Source, ProviderRepoName(provider.Name), version)
	}
	if h.ProgressUpdater != nil {
		h.ProgressUpdater.Update(provider.Name, ui.StatusInProgress, "Verifying...", 1)
	}
	l.Debug().Str("url", checksumsURL).Str("path", checksumsPath).Msg("downloading checksums file")
	// download checksums
	osFs := file.NewOsFs()
	if err := osFs.DownloadFile(ctx, checksumsPath, checksumsURL, nil); err != nil {
		l.Error().Err(err).Msg("failed to download checksums file")
		return false
	}
	l.Debug().Str("url", checksumsURL).Str("path", checksumsPath).Msg("downloading checksums signature")
	// download checksums signature
	if err := osFs.DownloadFile(ctx, checksumsPath+".sig", checksumsURL+".sig", nil); err != nil {
		l.Error().Err(err).Msg("failed to download signature file")
		return false
	}
	err := validateFile(checksumsPath, checksumsPath+".sig")
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

func (h Hub) downloadProvider(ctx context.Context, provider Provider, requestedVersion string, noVerify bool) (ProviderBinary, error) {

	if !h.verifyRegistered(provider.Source, provider.Name, requestedVersion, noVerify) {
		return ProviderBinary{}, fmt.Errorf("provider plugin %s@%s not registered at https://hub.cloudquery.io", provider.Name, requestedVersion)
	}
	// build fully qualified plugin directory for given plugin
	pluginDir := filepath.Join(h.PluginDirectory, provider.Source, provider.Name)
	osFs := file.NewOsFs()
	if err := osFs.MkdirAll(pluginDir, os.ModePerm); err != nil {
		return ProviderBinary{}, err
	}

	// Create a new progress updater callback func
	var progressCB ui.ProgressUpdateFunc
	if h.ProgressUpdater != nil {
		progressCB = ui.CreateProgressUpdater(h.ProgressUpdater, fmt.Sprintf("%s@%s", ProviderRepoName(provider.Name), requestedVersion))
	}

	providerURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", provider.Source, ProviderRepoName(provider.Name), requestedVersion, getPluginBinaryName(provider.Name))
	providerPath := h.getProviderPath(provider.Source, provider.Name, requestedVersion)
	if err := osFs.DownloadFile(ctx, providerPath, providerURL, progressCB); err != nil {
		return ProviderBinary{}, fmt.Errorf("plugin %s/%s@%s failed to download: %w", provider.Source, provider.Name, requestedVersion, err)
	}

	if ok := h.verifyProvider(ctx, provider, requestedVersion); !ok {
		return ProviderBinary{}, fmt.Errorf("plugin %s/%s@%s failed to verify", provider.Source, provider.Name, requestedVersion)
	}

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
		return "", fmt.Errorf("failed to find provider[%s] latest version", providerName)
	}
	return doc.Documents[0].Fields.Tag.Val, nil
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

func (h Hub) isProviderRegistered(org, provider string) bool {
	u := fmt.Sprintf(h.url, org, provider)
	res, err := http.Get(u)
	if err != nil {
		log.Error().Err(err).Msg("failed to check if provider is registered")
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

func GetBinarySuffix() string {
	var suffix = ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return fmt.Sprintf("%s_%s%s", runtime.GOOS, runtime.GOARCH, suffix)
}
