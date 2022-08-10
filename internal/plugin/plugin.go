// pluginmanager take care of lifecycle of plugins.
// Including: downloading, upgrading, spawning, closing
// Currently we use github releases as our plugin store. We might change in the future
// to our own hosted one.
package plugin

import (
	"archive/zip"
	"fmt"
	"io"
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

type Plugin struct {
	cmd          *exec.Cmd
	conn         *grpc.ClientConn
	sourceClient *clients.SourceClient
	// logger       zerolog.Logger
}

// func (p *Plugin) Write(bytes []byte) (n int, err error) {
// 	p.logger.
// 	// p.sourceClient
// 	return 0, nil
// }

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

// DownloadSource downloads a plugin from the specified registry and return the path to plugin
func (p *PluginManager) DownloadSource(ctx context.Context, spec specs.SourceSpec) (string, error) {
	switch spec.Registry {
	case "local", "grpc":
		fmt.Printf("Skipping plugin download. registry: %s, path: %s\n", spec.Registry, spec.Path)
		p.logger.Info().Str("registry", spec.Registry).Msg("Skiping plugin download")
		return "", nil
	case "github":
		return p.downloadSourceGitHub(ctx, spec)
	default:
		return "", errors.Errorf("unknown registry: %s", spec.Registry)
	}
}

func (p *PluginManager) downloadSourceGitHub(ctx context.Context, spec specs.SourceSpec) (string, error) {
	pathSplit := strings.Split(spec.Path, "/")
	org, repo := pathSplit[0], pathSplit[1]
	pluginName := fmt.Sprintf("cq-source-%s_%s_%s", repo, runtime.GOOS, runtime.GOARCH)
	dirPath := filepath.Join(p.directory, "plugins", spec.Registry, org, repo, spec.Version)
	pluginPath := filepath.Join(dirPath, pluginName)
	if _, err := os.Stat(pluginPath); err == nil {
		fmt.Printf("Plugin already exists at %s. Skipping download.\n", pluginPath)
		p.logger.Info().Str("path", pluginPath).Msg("Plugin already exists. Skipping download.")
		return "", nil
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", errors.Wrapf(err, "failed to create plugin directory: %s\n", dirPath)
	}
	// we use convention over configuration and we use github as our registry. Similar to how terraform and homebrew work.
	// For example:
	// https://github.com/cloudquery/cq-source-aws/releases/download/v1.0.1/cq-source-aws_darwin_amd64.zip
	pluginUrl := fmt.Sprintf("https://github.com/%s/cq-source-%s/releases/download/%s/cq-source-%s_%s_%s.zip", org, repo, spec.Version, repo, runtime.GOOS, runtime.GOARCH)
	if spec.Version == "latest" {
		// https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases
		pluginUrl = fmt.Sprintf("https://github.com/%s/cq-source-%s/releases/latest/download/cq-source-%s_%s_%s.zip", org, repo, repo, runtime.GOOS, runtime.GOARCH)
	}
	fmt.Printf("Downloading plugin from: %s to: %s.zip \n", pluginUrl, pluginPath)
	if err := downloadFile(pluginPath+".zip", pluginUrl); err != nil {
		return "", errors.Wrap(err, "failed to download plugin")
	}
	archive, err := zip.OpenReader(pluginPath + ".zip")
	if err != nil {
		return "", errors.Wrap(err, "failed to open plugin archive")
	}
	fileInArchive, err := archive.Open("cq-source-" + repo)
	if err != nil {
		return "", errors.Wrapf(err, "failed to open plugin archive: cq-source-%s", repo)
	}
	out, err := os.OpenFile(pluginPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create file: %s", pluginPath)
	}
	_, err = io.Copy(out, fileInArchive)
	if err != nil {
		return "", errors.Wrap(err, "failed to copy body to file")
	}
	return pluginPath, nil
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
	var grpcTarget string
	var pluginPath string
	switch spec.Registry {
	case "grpc":
		// This is a special case as we dont spawn any process
		conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to dial grpc target: %s", grpcTarget)
		}
		pl.conn = conn
		pl.sourceClient = clients.NewSourceClient(conn)
		p.plugins[spec.Registry][spec.Path] = &pl
		return p.plugins[spec.Registry][spec.Path].sourceClient, nil
	case "local":
		grpcTarget = unixSocketPrefix + spec.Path
		pluginPath = spec.Path
	case "github":
		var err error
		pluginPath, err = p.downloadSourceGitHub(ctx, spec)
		if err != nil {
			return nil, errors.Wrap(err, "failed to download plugin")
		}
		grpcTarget = unixSocketPrefix + spec.Path
	default:
		return nil, fmt.Errorf("unknown registry: %s", spec.Registry)
	}
	// spawn the plugin first and then connect
	if err := os.MkdirAll(filepath.Dir(grpcTarget), 0755); err != nil {
		return nil, errors.Wrapf(err, "failed to create unixpath directory: %s", filepath.Dir(grpcTarget))
	}
	cmd := exec.Command(pluginPath, "serve", "--network", "unix", "--address", grpcTarget,
		"--log-level", p.logger.GetLevel().String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = p.logger
	if err := cmd.Start(); err != nil {
		return nil, errors.Wrapf(err, "failed to start plugin: %s", pluginPath)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("plugin %s exited with error: %v\n", spec.Path, err)
			p.logger.Error().Err(err).Str("plugin", spec.Path).Msg("plugin exited")
		}
	}()
	conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("failed to kill plugin", err)
		}
		return nil, err
	}
	pl.cmd = cmd
	pl.conn = conn
	pl.sourceClient = clients.NewSourceClient(conn)
	p.plugins[spec.Registry][spec.Path] = &pl
	return p.plugins[spec.Registry][spec.Path].sourceClient, nil
}

func (p *PluginManager) CloseAll(ctx context.Context) error {
	for registryName, registry := range p.plugins {
		for path, pl := range registry {
			p.logger.Info().Str("registry", registryName).Str("plugin", path).Msg("closing connection to plugin")
			if err := pl.conn.Close(); err != nil {
				p.logger.Error().Str("registry", registryName).Str("plugin", path).Err(err)
			}
			pl.conn = nil
			if pl.cmd != nil && pl.cmd.Process != nil {
				if err := pl.cmd.Process.Kill(); err != nil {
					p.logger.Error().Str("registry", registryName).Str("plugin", path).Err(err)
				}
			}
		}
	}
	return nil
}
