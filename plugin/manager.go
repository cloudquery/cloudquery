package plugin

import (
	"fmt"
	"sync"
)

var (
	doOnce   sync.Once
	instance *Manager
)

// Manager handles CQProviders that can be either embedded (self-run provider for example) or remote using go_plugin
// Important: manager is currently not thread safe
type Manager struct {
	clients map[string]managedPlugin
}

// Shutdown closes all clients and cleans the managed clients
func (p *Manager) Shutdown() {
	for _, c := range p.clients {
		c.Close()
	}
	// create fresh map
	p.clients = make(map[string]managedPlugin)
}

func (p *Manager) GetProvider(providerName, version string) (CQProvider, error) {
	cq, ok := p.clients[providerName]
	if !ok {
		return nil, fmt.Errorf("plugin %s@%s does not exist", providerName, version)
	}
	return cq.Provider(), nil
}

func (p *Manager) AddEmbeddedPlugin(providerName string, cqp CQProvider) {
	p.clients[providerName] = newEmbeddedPlugin(providerName, "latest", cqp)
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

func (p *Manager) GetOrCreateProvider(providerName, version string) (CQProvider, error) {
	provider, err := p.GetProvider(providerName, version)
	if provider != nil || err == nil {
		return provider, err
	}
	// Create RPC client and initialize CQProvider
	return p.createProvider(providerName, version)
}

func (p *Manager) createProvider(providerName, version string) (CQProvider, error) {
	mPlugin, err := newRemotePlugin(providerName, version)
	if err != nil {
		return nil, err
	}
	p.clients[providerName] = mPlugin
	return mPlugin.Provider(), nil
}

func GetManager() *Manager {
	doOnce.Do(
		func() {
			instance = &Manager{
				clients: make(map[string]managedPlugin),
			}
		})
	return instance
}
