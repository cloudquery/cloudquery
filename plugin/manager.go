package plugin

import (
	"fmt"
	"github.com/cloudquery/cq-provider-sdk/proto"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Manager handles CQProviders that can be either embedded (self-run provider for example) or remote using go_plugin
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

func (p *Manager) GetProvider(providerName, version string) (proto.CQProvider, error) {
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

func (p *Manager) GetOrCreateProvider(providerName, version string) (proto.CQProvider, error) {
	provider, err := p.GetProvider(providerName, version)
	if provider != nil || err == nil {
		return provider, err
	}
	// Create RPC client and initialize CQProvider
	return p.createProvider(providerName, version)
}

func (p *Manager) createProvider(providerName, version string) (proto.CQProvider, error) {
	mPlugin, err := newRemotePlugin(providerName, version)
	if err != nil {
		return nil, err
	}
	p.clients[providerName] = mPlugin
	return mPlugin.Provider(), nil
}
