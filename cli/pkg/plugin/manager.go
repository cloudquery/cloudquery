package plugin

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/cli/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type CreationOptions struct {
	Provider registry.Provider
	// Alias to name plugin on creation
	Alias string
	// Environment variables to pass to plugin binary on creation
	Env []string
}

// Manager handles lifecycle execution of CloudQuery providers
type Manager struct {
	// whether manager allows executing plugins in reattach mode or not, by default reattach is disabled.
	allowReattach bool
	// registry allows access to download of providers from a remote source
	registry registry.Registry
	// clients is a map of all plugins created and managed by the manager
	clients Plugins
}

type ManagerOption func(m *Manager)

// WithAllowReattach allows plugin reattach to be supported by Manager
func WithAllowReattach() ManagerOption {
	return func(m *Manager) {
		m.allowReattach = true
	}
}

func NewManager(r registry.Registry, opts ...ManagerOption) (*Manager, error) {
	m := &Manager{
		clients:       make(map[string]Plugin),
		registry:      r,
		allowReattach: false,
	}
	// apply options
	for _, o := range opts {
		o(m)
	}

	if m.allowReattach {
		if err := m.reattachProviders(); err != nil {
			return nil, err
		}
	}

	return m, nil
}

// DownloadProviders downloads one or more registry.Provider from the registry.Registry, if we want to skip
// provider verification when downloading pass true on `noVerify`
func (m *Manager) DownloadProviders(ctx context.Context, providers []registry.Provider, noVerify bool) ([]registry.ProviderBinary, error) {
	log.Debug().Interface("providers", providers).Msg("Downloading required providers")
	downloaded := make([]registry.ProviderBinary, len(providers))
	for i, rp := range providers {
		if m.IsReattachProvider(rp) {
			ui.ColorizedOutput(ui.ColorInfo, fmt.Sprintf("Skipping provider %s download, using reattach instead\n", ui.Colorize(ui.ColorSuccessBold, false, rp.Name)))
			continue
		}
		log.Info().Str("name", rp.Name).Str("version", rp.Version).Msg("Downloading provider")
		details, err := m.registry.Download(ctx, rp, noVerify)
		if err != nil {
			return nil, err
		}
		downloaded[i] = details
	}
	return downloaded, nil
}

// CreatePlugin creates a plugin based on CreationOptions
func (m *Manager) CreatePlugin(opts *CreationOptions) (Plugin, error) {
	_, providerName, err := registry.ParseProviderName(opts.Provider.Name)
	if err != nil {
		return nil, err
	}

	p := m.clients.Get(opts.Provider, opts.Alias)
	if p != nil {
		log.Debug().Stringer("provider", opts.Provider).Str("alias", opts.Alias).Msg("using existing plugin")
		return p, nil
	}
	log.Info().Str("provider", providerName).Str("alias", opts.Alias).Msg("plugin doesn't exist, creating...")
	details, err := m.registry.Get(opts.Provider.Name, opts.Provider.Version)
	if err != nil {
		return nil, fmt.Errorf("no such provider %s. plugin might be missing from directory or wasn't downloaded", providerName)
	}
	p, err = m.createProvider(&details, opts.Alias, opts.Env)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Shutdown closes all clients and cleans the managed clients
func (m *Manager) Shutdown() {
	for _, c := range m.clients {
		m.ClosePlugin(c)
	}
	// create fresh map
	m.clients = make(map[string]Plugin)
}

// TODO: support Closer from plugin rather then calling manager

// ClosePlugin kills a plugin instance and removes it from the managed plugins.
func (m *Manager) ClosePlugin(p Plugin) {
	if p.Version() == Unmanaged {
		log.Warn().Str("provider", p.Name()).Msg("not closing unmanaged provider")
		return
	}
	if err := m.killProvider(p.Name()); err != nil {
		log.Warn().Str("provider", p.Name()).Msg("failed to kill provider")
	}
}

func (m *Manager) IsReattachProvider(rp registry.Provider) bool {
	_, ok := m.clients[rp.Name]
	return ok
}

func (m *Manager) reattachProviders() error {
	// used primarily by the SDK's acceptance testing framework.
	unmanagedProviders, err := serve.ParseReattachProviders(viper.GetString("reattach-providers"))
	if err != nil {
		return err
	}
	for name, cfg := range unmanagedProviders {
		log.Debug().Str("name", name).Str("address", cfg.Addr.String()).Int("pid", cfg.Pid).Msg("reattaching unmanaged plugin")
		plugin, err := newUnmanagedPlugin(name, cfg)
		if err != nil {
			return err
		}
		m.clients[name] = plugin
	}
	return nil
}

func (m *Manager) killProvider(providerName string) error {
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

func (m *Manager) createProvider(details *registry.ProviderBinary, alias string, env []string) (Plugin, error) {
	mPlugin, err := newRemotePlugin(details, alias, env)
	if err != nil {
		return nil, err
	}
	m.clients[mPlugin.Name()] = mPlugin
	return mPlugin, nil
}
