// pluginmanager take care of lifecycle of plugins.
// Including: downloading, upgrading, spawning, closing
// Currently we use github releases as our plugin store. We might change in the future
// to our own hosted one.
package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"context"

	"github.com/cloudquery/cloudquery/internal/destinations"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const unixSocketPrefix = "/tmp/cq-plugins/"

func hasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// get unnormalized plugin name and returns normalized/
// aws -> cloudquery, aws, latest
// cloudquery/aws -> cloudquery, aws, latest
// aws@v1.0.0 -> cloudquery, aws, v1.0.0
func parsePluginName(name string) (string, string, string, error) {
	if len(name) == 0 {
		return "", "", "", errors.New("plugin name is empty")
	}

	organization := "cloudquery"
	pluginName := name
	version := "latest"

	pluginPart := strings.Split(name, "@")
	if len(pluginPart) > 2 {
		return "", "", "", errors.Errorf("invalid plugin name: %s. only one @ is allowed", name)
	}
	if len(pluginPart) == 2 {
		version = pluginPart[1]
	}

	pluginPart = strings.Split(pluginPart[0], "/")
	if len(pluginPart) > 2 {
		return "", "", "", errors.Errorf("invalid plugin name: %s. only one / is allowed", name)
	}

	if len(pluginPart) == 2 {
		organization = pluginPart[0]
	}

	return organization, pluginName, version, nil
}

type Plugin struct {
	cmd          *exec.Cmd
	conn         *grpc.ClientConn
	sourceClient *clients.SourceClient
}

type PluginManager struct {
	// plugins      map[string]*exec.Cmd
	// sourceClient map[string]*clients.SourceClient
	plugins   map[string]map[string]*Plugin
	logger    zerolog.Logger
	directory string
}

type PluginManagerOption func(*PluginManager)

func WithLogger(logger zerolog.Logger) func(*PluginManager) {
	return func(p *PluginManager) {
		p.logger = logger
	}
}

func WithDirectory(directory string) func(*PluginManager) {
	return func(p *PluginManager) {
		p.directory = directory
	}
}

func NewPluginManager(opts ...PluginManagerOption) *PluginManager {
	p := &PluginManager{
		logger:    log.Logger,
		directory: "./.cq",
		plugins:   make(map[string]map[string]*Plugin),
	}
	// initialize all plugins registry
	p.plugins["local"] = make(map[string]*Plugin)
	p.plugins["github"] = make(map[string]*Plugin)
	p.plugins["grpc"] = make(map[string]*Plugin)
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *PluginManager) Download(ctx context.Context, spec specs.SourceSpec) error {
	if spec.Registry == "local" || spec.Registry == "grpc" {
		p.logger.Info().Str("registry", spec.Registry).Msg("skiping plugin download")
		return nil
	}

	pathSplit := strings.Split(spec.Path, "/")
	org, repo := pathSplit[0], pathSplit[1]
	pluginName := fmt.Sprintf("cq-provider-%s_%s_%s", repo, runtime.GOOS, runtime.GOARCH)
	dirPath := filepath.Join(p.directory, "plugins", spec.Registry, org, repo, spec.Version)
	pluginPath := filepath.Join(dirPath, pluginName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return errors.Wrapf(err, "failed to create plugin directory: %s", dirPath)
	}
	if _, err := os.Stat(pluginPath); err == nil {
		fmt.Printf("Plugin already exists at %s. Skipping download.\n", pluginPath)
		p.logger.Info().Str("path", pluginPath).Msg("Plugin already exists. Skipping download.")
		return nil
	}

	// we use convention over configuration and we use github as our registry. Similar to how terraform and homebrew work.
	// For example:
	// https://github.com/cloudquery/cq-provider-aws/releases/download/v1.0.1/cq-provider-aws_darwin_amd64.zip
	pluginUrl := fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/download/%s/cq-provider-%s_%s_%s", org, repo, spec.Version, repo, runtime.GOOS, runtime.GOARCH)
	if err := downloadFile(pluginPath, pluginUrl); err != nil {
		return errors.Wrap(err, "failed to download plugin")
	}

	return nil
}

func (p *PluginManager) GetDestinationClient(ctx context.Context, spec specs.DestinationSpec, opts plugins.DestinationPluginOptions) (*clients.DestinationClient, error) {
	switch spec.Name {
	case "postgresql":
		return clients.NewLocalDestinationClient(&destinations.PostgreSqlPlugin{}), nil
	default:
		return nil, errors.Errorf("unknown destination type: %s", spec.Name)
	}
}

func (p *PluginManager) GetSourcePluginClient(ctx context.Context, spec specs.SourceSpec) (*clients.SourceClient, error) {
	if p.plugins[spec.Registry] != nil && p.plugins[spec.Registry][spec.Path] != nil {
		return p.plugins[spec.Registry][spec.Path].sourceClient, nil
	}
	pl := Plugin{}
	grpcTarget := spec.Path
	switch spec.Registry {
	case "local":
		grpcTarget := unixSocketPrefix + spec.Path
		cmd := exec.Command(spec.Path, "serve", "--network", "unix", "--address", grpcTarget)
		if err := cmd.Start(); err != nil {
			return nil, errors.Wrap(err, "failed to start plugin")
		}
		pl.cmd = cmd
	case "github":
		if err := p.Download(ctx, spec); err != nil {
			return nil, errors.Wrap(err, "failed to download plugin")
		}
		pathSplit := strings.Split(spec.Path, "/")
		org, repo := pathSplit[0], pathSplit[1]
		pluginName := fmt.Sprintf("cq-provider-%s_%s_%s", repo, runtime.GOOS, runtime.GOARCH)
		pluginPath := filepath.Join(p.directory, "plugins/github", org, repo, spec.Version, pluginName)
		grpcTarget := unixSocketPrefix + spec.Path
		cmd := exec.Command(pluginPath, "serve", "--network", "unix", "--address", grpcTarget)
		if err := cmd.Start(); err != nil {
			return nil, errors.Wrap(err, "failed to start plugin")
		}
		pl.cmd = cmd
	default:
		return nil, fmt.Errorf("unknown registry: %s", spec.Registry)
	}
	conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		pl.cmd.Process.Kill()
		return nil, err
	}
	pl.conn = conn
	pl.sourceClient = clients.NewSourceClient(conn)
	p.plugins[spec.Registry][spec.Path] = &pl
	return p.plugins[spec.Registry][spec.Path].sourceClient, nil
}

func (p *PluginManager) CloseAll(ctx context.Context) error {
	for registryName, registry := range p.plugins {
		for path, pl := range registry {
			p.logger.Info().Str("registry", registryName).Str("plugin", path).Msg("closing plugin")
			if err := pl.conn.Close(); err != nil {
				p.logger.Error().Str("registry", registryName).Str("plugin", path).Err(err)
			}
			if err := pl.cmd.Process.Kill(); err != nil {
				p.logger.Error().Str("registry", registryName).Str("plugin", path).Err(err)
			}
		}
	}
	return nil
}
