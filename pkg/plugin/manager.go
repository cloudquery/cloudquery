package plugin

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Manager handles lifecycle execution of CloudQuery providers
type Manager struct {
	clients map[string]Plugin
}

func NewManager() (*Manager, error) {
	// primarily by the SDK's acceptance testing framework.
	unmanagedProviders, err := serve.ParseReattachProviders(viper.GetString("reattach-providers"))
	if err != nil {
		return nil, err
	}

	clients := make(map[string]Plugin)
	for name, config := range unmanagedProviders {
		log.Debug().Str("name", name).Str("address", config.Addr.String()).Int("pid", config.Pid).Msg("reattaching unmanaged plugin")
		plugin, err := newUnmanagedPlugin(name, config)
		if err != nil {
			return nil, err
		}
		clients[name] = plugin
	}
	return &Manager{
		clients: clients,
	}, nil
}

// Shutdown closes all clients and cleans the managed clients
func (p *Manager) Shutdown() {
	for _, c := range p.clients {
		c.Close()
	}
	// create fresh map
	p.clients = make(map[string]Plugin)
}

func (p *Manager) GetProvider(providerName, version string) (cqproto.CQProvider, error) {
	cq, ok := p.clients[providerName]
	if !ok {
		return nil, fmt.Errorf("plugin %s@%s does not exist", providerName, version)
	}
	return cq.Provider(), nil
}

func (p *Manager) KillProvider(providerName string) error {

	client, ok := p.clients[providerName]
	if !ok {
		return fmt.Errorf("client for provider %s does not exist", providerName)
	}
	client.Close()
	delete(p.clients, providerName)
	return nil
}

func (p *Manager) GetOrCreateProvider(details *registry.ProviderDetails) (cqproto.CQProvider, error) {
	provider, err := p.GetProvider(details.Name, details.Version)
	if provider != nil || err == nil {
		return provider, err
	}
	// Create RPC client and initialize CQProvider
	return p.createProvider(details)
}

func (p *Manager) createProvider(details *registry.ProviderDetails) (cqproto.CQProvider, error) {
	mPlugin, err := newRemotePlugin(details)
	if err != nil {
		return nil, err
	}
	p.clients[details.Name] = mPlugin
	return mPlugin.Provider(), nil
}

func (p *Manager) ListUnmanaged() map[string]registry.ProviderDetails {
	unmanged := make(map[string]registry.ProviderDetails)
	for k, v := range p.clients {
		if _, ok := v.(*unmanagedPlugin); !ok {
			continue
		}
		unmanged[k] = registry.ProviderDetails{
			Name:    v.Name(),
			Version: v.Version(),
		}
	}
	return unmanged
}
