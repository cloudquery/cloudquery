package client

import (
	"fmt"
	"github.com/cloudquery/cq-provider-sdk/proto"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cloudquery/cloudquery/logging"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const defaultOrganization = "cloudquery"

type Plugin interface {
	Name() string
	Version() string
	Provider() proto.CQProvider
	Close()
}

type managedPlugin struct {
	name     string
	version  string
	client   *plugin.Client
	provider proto.CQProvider
}

// NewRemotePlugin creates a new remoted plugin using go_plugin
func newRemotePlugin(providerName, version string) (*managedPlugin, error) {
	pluginPath, _ := GetProviderPath(providerName, version)
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: serve.Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: serve.PluginMap,
		},
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		SyncStderr:       os.Stderr,
		SyncStdout:       os.Stdout,
		Logger:           logging.NewZHcLog(&log.Logger, ""),
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

	provider, ok := raw.(proto.CQProvider)
	if !ok {
		client.Kill()
		return nil, fmt.Errorf("failed to cast plugin")
	}
	return &managedPlugin{
		name:     providerName,
		version:  version,
		client:   client,
		provider: provider,
	}, nil
}

func (m managedPlugin) Name() string { return m.name }

func (m managedPlugin) Version() string { return m.version }

func (m managedPlugin) Provider() proto.CQProvider { return m.provider }

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
	provider proto.CQProvider
}

// newUnmanagedPlugin attaches to and existing running plugin  a new unmanaged plugin using go_plugin
func newUnmanagedPlugin(providerName string, config *plugin.ReattachConfig) (*unmanagedPlugin, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  serve.Handshake,
		Plugins:          serve.PluginMap,
		Reattach:         config,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		SyncStderr:       os.Stderr,
		SyncStdout:       os.Stdout,
		Logger:           logging.NewZHcLog(&log.Logger, ""),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		return nil, err
	}

	provider, ok := raw.(proto.CQProvider)
	if !ok {
		return nil, fmt.Errorf("failed to cast plugin")
	}
	return &unmanagedPlugin{
		name:     providerName,
		config:   config,
		client:   client,
		provider: provider,
	}, nil
}

func (m unmanagedPlugin) Name() string { return m.name }

func (m unmanagedPlugin) Version() string { return "testing" }

func (m unmanagedPlugin) Provider() proto.CQProvider { return m.provider }

func (m unmanagedPlugin) Close() {}

// GetProviderPath returns expected path of provider on file system from name and version of plugin
func GetProviderPath(name string, version string) (string, error) {
	org := defaultOrganization
	split := strings.Split(name, "/")
	if len(split) == 2 {
		org = split[0]
		name = split[1]
	}

	pluginDir := viper.GetString("plugin-dir")

	extension := ""
	if runtime.GOOS == "windows" {
		extension = ".exe"
	}
	return filepath.Join(pluginDir, ".cq", "providers", org, name, fmt.Sprintf("%s-%s-%s%s", version, runtime.GOOS, runtime.GOARCH, extension)), nil
}
