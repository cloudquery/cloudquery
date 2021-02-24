package plugin

import (
	"fmt"
	"sync"
)

var (
	doOnce   sync.Once
	instance *Manager
)

type Manager struct {
	clients map[string]ManagedPlugin
}

func (p *Manager) Shutdown() {
	for _, c := range p.clients {
		c.Close()
	}
	// create fresh map
	p.clients = make(map[string]ManagedPlugin)
}

func (p *Manager) GetProvider(providerName, version string) (CQProvider, error) {
	cq, ok := p.clients[providerName]
	if !ok {
		return nil, fmt.Errorf("plugin %s@%s does not exist", providerName, version)
	}
	return cq.Provider(), nil
}

func (p *Manager) AddEmbeddedPlugin(providerName string, cqp CQProvider) {
	p.clients[providerName] = NewEmbeddedPlugin(providerName, "latest", cqp)
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
	mPlugin, err := NewRemotePlugin(providerName, version)
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
				clients: make(map[string]ManagedPlugin),
			}
		})
	return instance
}
