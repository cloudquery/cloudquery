package plugin

import (
	"fmt"
	"github.com/cloudquery/cloudquery/logging"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultOrganization = "cloudquery"

type ManagedPlugin interface {
	Name() string
	Version() string
	Provider() CQProvider
	Close()
}

type RemotePlugin struct {
	name     string
	version  string
	client   *plugin.Client
	provider CQProvider
}

type EmbeddedPlugin struct {
	name     string
	version  string
	provider CQProvider
}

func NewRemotePlugin(providerName, version string) (*RemotePlugin, error) {
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

	provider, ok := raw.(CQProvider)
	if !ok {
		client.Kill()
		return nil, fmt.Errorf("failed to cast plugin")
	}
	return &RemotePlugin{
		name:     providerName,
		version:  version,
		client:   client,
		provider: provider,
	}, nil
}

func NewEmbeddedPlugin(providerName, version string, p CQProvider) *EmbeddedPlugin {
	return &EmbeddedPlugin{
		name:     providerName,
		version:  version,
		provider: p,
	}
}

func (e EmbeddedPlugin) Name() string { return e.name }

func (e EmbeddedPlugin) Version() string { return e.version }

func (e EmbeddedPlugin) Provider() CQProvider { return e.provider }

func (e EmbeddedPlugin) Close() { return }

func (r RemotePlugin) Name() string { return r.name }

func (r RemotePlugin) Version() string { return r.version }

func (r RemotePlugin) Provider() CQProvider { return r.provider }

func (r RemotePlugin) Close() {
	if r.client == nil {
		return
	}
	r.client.Kill()
}

func getProviderPath(name string, version string) (string, error) {
	org := defaultOrganization
	split := strings.Split(name, "/")
	if len(split) == 2 {
		org = split[0]
		name = split[1]
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	extension := ""
	if runtime.GOOS == "windows" {
		extension = ".exe"
	}
	return filepath.Join(workingDir, ".cq", "providers", org, name, fmt.Sprintf("%s-%s-%s%s", version, runtime.GOOS, runtime.GOARCH, extension)), nil
}
