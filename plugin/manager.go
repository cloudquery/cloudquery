package plugin

import (
	"fmt"
	"github.com/cloudquery/cloudquery/logging"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"sync"
)

var (
	doOnce   sync.Once
	instance *Manager
)

type Manager struct {
	activeClients map[string]*plugin.Client
	registry      map[string]CQProvider

	lock sync.RWMutex
}

func (p *Manager) Setup() error {
	return nil
}

func (p *Manager) Shutdown() {
	p.lock.Lock()
	defer p.lock.Unlock()
	for n, c := range p.activeClients {
		c.Kill()
		delete(p.activeClients, n)
	}
	p.registry = make(map[string]CQProvider)
}


func (p *Manager) GetProvider(providerName, version string) (CQProvider, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	cq, ok := p.registry[providerName]
	if !ok {
		return nil, fmt.Errorf("plugin %s@%s does not exist", providerName, version)
	}
	return cq, nil
}

func (p *Manager) SetProvider(providerName string, cqp CQProvider) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.registry[providerName] = cqp
}

func (p *Manager) KillProvider(providerName string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	client, ok := p.activeClients[providerName]
	if !ok {
		return fmt.Errorf("client for provider %s does not exist", providerName)
	}
	client.Kill()
	delete(p.activeClients, providerName)
	return nil
}

func (p *Manager) GetOrCreateProvider(providerName, version string) (CQProvider, error) {

	provider, err := p.GetProvider(providerName, version)
	if provider != nil && err == nil {
		return provider, err
	}
	// Create RPC client and initialize CQProvider
	return p.createProvider(providerName, version)
}


func (p *Manager) createProvider(providerName, version string) (CQProvider, error) {
	pluginPath, _ := getProviderPath(providerName, version)
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: PluginMap,
		},
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		SyncStderr:       os.Stderr,
		SyncStdout:       os.Stdout,
		Logger: logging.NewZHcLog(&log.Logger, ""),
	})
	rpcClient, err := client.Client()
	if err != nil {
		client.Kill()
		return nil, err
	}
	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		client.Kill()
		return nil, err
	}

	provider, ok := raw.(CQProvider)
	if !ok {
		client.Kill()
		return nil, fmt.Errorf("failed to cast plugin")
	}
	p.activeClients[providerName] = client
	p.registry[providerName] = provider
	return provider, nil
}

func GetManager() *Manager {
	doOnce.Do(
		func() {
			instance = &Manager{
				activeClients: make(map[string]*plugin.Client),
				registry:      make(map[string]CQProvider),
				lock:          sync.RWMutex{},
			}
		})
	return instance
}
