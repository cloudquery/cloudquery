package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/serve"

	"github.com/hashicorp/go-plugin"
	zerolog "github.com/rs/zerolog/log"
)

const (
	Unmanaged = "unmanaged"
)

// PluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"provider": &cqproto.CQPlugin{},
}

type Plugin interface {
	Name() string
	Version() string
	ProtocolVersion() int
	Provider() cqproto.CQProvider
	Close()
}

type managedPlugin struct {
	name     string
	version  string
	client   *plugin.Client
	provider cqproto.CQProvider
}

// NewRemotePlugin creates a new remoted plugin using go_plugin
func newRemotePlugin(details *registry.ProviderDetails, alias string, env []string) (*managedPlugin, error) {
	cmd := exec.Command(details.FilePath)
	cmd.Env = append(cmd.Env, env...)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: serve.Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			cqproto.V4: pluginMap,
		},
		Managed:          true,
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           logging.NewZHcLog(&zerolog.Logger, ""),
	})
	rpcClient, err := client.Client()
	if err != nil {
		client.Kill()
		// give a more clear message to the user
		if strings.Contains(err.Error(), "Incompatible API version") {
			return nil, fmt.Errorf("%w. Please upgrade to a latest version of this provider", err)
		}
		return nil, err
	}
	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		client.Kill()
		return nil, err
	}

	provider, ok := raw.(cqproto.CQProvider)
	if !ok {
		client.Kill()
		return nil, fmt.Errorf("failed to cast plugin")
	}
	name := details.Name
	if alias != "" {
		name = fmt.Sprintf("%s_%s", name, alias)
	}
	return &managedPlugin{
		name:     name,
		version:  details.Version,
		client:   client,
		provider: provider,
	}, nil
}

func (m managedPlugin) Name() string { return m.name }

func (m managedPlugin) Version() string { return m.version }

func (m managedPlugin) ProtocolVersion() int { return m.client.NegotiatedVersion() }

func (m managedPlugin) Provider() cqproto.CQProvider { return m.provider }

func (m managedPlugin) Close() {
	if m.client == nil {
		return
	}
	m.client.Kill()
}

type unmanagedPlugin struct {
	name     string
	config   *plugin.ReattachConfig
	client   *plugin.Client
	provider cqproto.CQProvider
}

// newUnmanagedPlugin attaches to and existing running plugin  a new unmanaged plugin using go_plugin
func newUnmanagedPlugin(providerName string, config *plugin.ReattachConfig) (*unmanagedPlugin, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  serve.Handshake,
		Plugins:          pluginMap,
		Reattach:         config,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		SyncStderr:       os.Stderr,
		SyncStdout:       os.Stdout,
		Logger:           logging.NewZHcLog(&zerolog.Logger, ""),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		return nil, err
	}

	provider, ok := raw.(cqproto.CQProvider)
	if !ok {
		return nil, fmt.Errorf("failed to cast plugin")
	}
	return &unmanagedPlugin{
		name:     fmt.Sprintf("%s_%s", providerName, providerName),
		config:   config,
		client:   client,
		provider: provider,
	}, nil
}

func (m unmanagedPlugin) Name() string { return m.name }

func (m unmanagedPlugin) Version() string { return Unmanaged }

func (m unmanagedPlugin) ProtocolVersion() int { return cqproto.Vunmanaged }

func (m unmanagedPlugin) Provider() cqproto.CQProvider { return m.provider }

func (m unmanagedPlugin) Close() {}
