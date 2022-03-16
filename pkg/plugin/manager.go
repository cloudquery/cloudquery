package plugin

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/go-hclog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

// Manager handles lifecycle execution of CloudQuery providers
type Manager struct {
	hub       *registry.Hub
	clients   map[string]Plugin
	providers map[string]registry.ProviderDetails
	logger    hclog.Logger
}

func NewManager(logger hclog.Logger, pluginDirectory string, registryURL string, updater ui.Progress) (*Manager, error) {
	// primarily by the SDK's acceptance testing framework.
	unmanagedProviders, err := serve.ParseReattachProviders(viper.GetString("reattach-providers"))
	if err != nil {
		return nil, err
	}
	clients := make(map[string]Plugin)
	for name, cfg := range unmanagedProviders {
		log.Debug().Str("name", name).Str("address", cfg.Addr.String()).Int("pid", cfg.Pid).Msg("reattaching unmanaged plugin")
		plugin, err := newUnmanagedPlugin(name, cfg)
		if err != nil {
			return nil, err
		}
		clients[name] = plugin
	}
	return &Manager{
		clients:   clients,
		logger:    logger,
		providers: make(map[string]registry.ProviderDetails),
		hub: registry.NewRegistryHub(registryURL, func(h *registry.Hub) {
			h.ProgressUpdater = updater
			h.PluginDirectory = pluginDirectory
		}),
	}, nil
}

// LoadExisting loads existing providers that are found by the hub in ProviderDirectory
func (m *Manager) LoadExisting(providers []*config.RequiredProvider) {
	for _, p := range providers {
		pd, err := m.hub.GetProvider(p.Name, p.Version)
		if err != nil {
			continue
		}
		m.providers[pd.Name] = pd
	}
}

func (m *Manager) DownloadProviders(ctx context.Context, providers []*config.RequiredProvider, noVerify bool) error {
	m.logger.Debug("Downloading required providers", "providers", providers)
	traceData := make([]string, len(providers))
	for i, rp := range providers {
		if _, ok := m.clients[rp.Name]; ok {
			m.logger.Debug("Skipping provider download, using reattach instead", "name", rp.Name, "version", rp.Version)
			traceData[i] = rp.Name + "@debug"
			continue
		}
		m.logger.Info("Downloading provider", "name", rp.Name, "version", rp.Version)
		details, err := m.hub.DownloadProvider(ctx, rp, noVerify)
		if err != nil {
			return err
		}
		m.providers[rp.Name] = details
		traceData[i] = rp.Name + "@" + details.Version
	}

	sort.Strings(traceData)
	trace.SpanFromContext(ctx).SetAttributes(
		attribute.StringSlice("providers", traceData),
	)

	return nil
}

func (m *Manager) CreatePlugin(providerName, alias string, env []string) (Plugin, error) {
	_, providerName, err := registry.ParseProviderName(providerName)
	if err != nil {
		return nil, err
	}
	p, ok := m.clients[providerName]
	if ok {
		return p, nil
	}
	m.logger.Info("plugin doesn't exist, creating...", "provider", providerName, "alias", alias)
	details, ok := m.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("no such provider %s. plugin might be missing from directory or wasn't downloaded", providerName)
	}
	p, err = m.createProvider(&details, alias, env)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *Manager) GetPluginDetails(providerName string) (registry.ProviderDetails, error) {
	details, ok := m.providers[providerName]
	if !ok {
		return registry.ProviderDetails{}, fmt.Errorf("provider %s doesn't exist", providerName)
	}
	return details, nil
}

// Shutdown closes all clients and cleans the managed clients
func (m *Manager) Shutdown() {
	for _, c := range m.clients {
		c.Close()
	}
	// create fresh map
	m.clients = make(map[string]Plugin)
}

func (m *Manager) KillProvider(providerName string) error {
	_, providerName, err := registry.ParseProviderName(providerName)
	if err != nil {
		return err
	}

	client, ok := m.clients[providerName]
	if !ok {
		return fmt.Errorf("client for provider %s does not exist", providerName)
	}
	client.Close()
	delete(m.clients, providerName)
	return nil
}

func (m *Manager) createProvider(details *registry.ProviderDetails, alias string, env []string) (Plugin, error) {
	mPlugin, err := newRemotePlugin(details, alias, env)
	if err != nil {
		return nil, err
	}
	m.clients[mPlugin.Name()] = mPlugin
	return mPlugin, nil
}
