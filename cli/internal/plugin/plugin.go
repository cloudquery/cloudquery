// pluginmanager take care of lifecycle of plugins.
// Including: downloading, upgrading, spawning, closing
// Currently we use github releases as our plugin store. We might change in the future
// to our own hosted one.
package plugin

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"context"

	"github.com/cloudquery/cloudquery/cli/internal/destinations/postgresql"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DestinationPlugin struct {
	cmd    *exec.Cmd
	conn   *grpc.ClientConn
	client *clients.DestinationClient
}

func (p *DestinationPlugin) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	if p.cmd != nil && p.cmd.Process != nil {
		p.cmd.Process.Kill()
	}
	return nil
}

func (p *DestinationPlugin) GetClient() *clients.DestinationClient {
	return p.client
}

type SourcePlugin struct {
	cmd    *exec.Cmd
	conn   *grpc.ClientConn
	client *clients.SourceClient
}

func (p *SourcePlugin) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	if p.cmd != nil && p.cmd.Process != nil {
		p.cmd.Process.Kill()
	}
	return nil
}

func (p *SourcePlugin) GetClient() *clients.SourceClient {
	return p.client
}

type PluginManager struct {
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
	}
	// initialize all plugins registry
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// DownloadSource downloads a plugin from the specified registry and return the path to plugin
func (p *PluginManager) DownloadSource(ctx context.Context, spec specs.Source) (string, error) {
	switch spec.Registry {
	case specs.RegistryLocal, specs.RegistryGrpc:
		fmt.Printf("Skipping plugin download. registry: %s, path: %s\n", spec.Registry, spec.Path)
		p.logger.Info().Str("registry", spec.Registry.String()).Msg("Skiping plugin download")
		return "", nil
	case specs.RegistryGithub:
		return p.downloadSourceGitHub(ctx, spec)
	default:
		return "", errors.Errorf("unknown registry: %s", spec.Registry)
	}
}

func (p *PluginManager) downloadSourceGitHub(ctx context.Context, spec specs.Source) (string, error) {
	pathSplit := strings.Split(spec.Path, "/")
	org, repo := pathSplit[0], pathSplit[1]
	pluginName := fmt.Sprintf("cq-source-%s_%s_%s", repo, runtime.GOOS, runtime.GOARCH)
	dirPath := filepath.Join(p.directory, "plugins", spec.Registry.String(), org, repo, spec.Version)
	pluginPath := filepath.Join(dirPath, pluginName)
	if _, err := os.Stat(pluginPath); err == nil {
		fmt.Printf("Plugin already exists at %s. Skipping download.\n", pluginPath)
		p.logger.Info().Str("path", pluginPath).Msg("Plugin already exists. Skipping download.")
		return pluginPath, nil
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

// NewDestination Plugin downloads the plugin, spanws the process (if needed)
// and return a new client. The calee is responsible for closing the plugin.
func (p *PluginManager) NewDestinationPlugin(ctx context.Context, spec specs.Destination) (*DestinationPlugin, error) {
	pl := DestinationPlugin{}
	// some destination plugins are compiled in for simplicity (so no need to download them and spawn grpc server)
	switch spec.Name {
	case "postgresql":
		pl.client = clients.NewLocalDestinationClient(postgresql.NewClient(p.logger))
		return &pl, nil
	default:
		return nil, fmt.Errorf("unknown destination plugin: %s", spec.Name)
	}
}

func (p *PluginManager) NewSourcePlugin(ctx context.Context, spec specs.Source) (*SourcePlugin, error) {
	pl := SourcePlugin{}
	// var grpcTarget string
	var pluginPath string
	switch spec.Registry {
	case specs.RegistryGrpc:
		// This is a special case as we dont spawn any process
		conn, err := grpc.Dial(spec.Path, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to dial grpc target: %s", spec.Path)
		}
		pl.conn = conn
		pl.client = clients.NewSourceClient(conn)
		return &pl, nil
	case specs.RegistryLocal:
		pluginPath = spec.Path
	case specs.RegistryGithub:
		var err error
		pluginPath, err = p.downloadSourceGitHub(ctx, spec)
		if err != nil {
			return nil, err
		}
		// grpcTarget = unixSocketPrefix + spec.Path
	default:
		return nil, fmt.Errorf("unknown registry: %s", spec.Registry)
	}
	grpcTarget := generateRandomUnixSocketName()
	// spawn the plugin first and then connect
	// if err := os.MkdirAll(filepath.Dir(grpcTarget), 0755); err != nil {
	// 	return nil, errors.Wrapf(err, "failed to create unixpath directory: %s", filepath.Dir(grpcTarget))
	// }
	cmd := exec.Command(pluginPath, "serve", "--network", "unix", "--address", grpcTarget,
		"--log-level", p.logger.GetLevel().String(), "--log-format", "json")
	reader, writer := io.Pipe()
	cmd.Stdout = writer
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, errors.Wrapf(err, "failed to start plugin: %s", pluginPath)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("plugin %s exited with error: %v\n", spec.Path, err)
			p.logger.Error().Err(err).Str("plugin", spec.Path).Msg("plugin exited")
		}
	}()
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			var structuredLogLine map[string]interface{}
			b := scanner.Bytes()
			if err := json.Unmarshal(b, &structuredLogLine); err != nil {
				p.logger.Err(err).Str("line", string(b)).Msg("failed to unmarshal log line from plugin")
			} else {
				jsonToLog(structuredLogLine, p.logger)
				// p.logger.js
				// p.logger.Output()
			}
		}
	}()
	// remove the socket file if it exists
	// os.Remove(grpcTarget)
	conn, err := grpc.Dial("unix://"+grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("failed to kill plugin", err)
		}
		return nil, err
	}
	pl.cmd = cmd
	pl.conn = conn
	pl.client = clients.NewSourceClient(conn)
	return &pl, nil
}
