// pluginmanager take care of lifecycle of plugins.
// Including: downloading, upgrading, spawning, closing
// Currently we use github releases as our plugin store. We might change in the future
// to our own hosted one.
package plugins

import (
	"archive/zip"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/destinations/postgresql"
	"github.com/cloudquery/cloudquery/cli/internal/versions"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SourcePlugin struct {
	cmd      *exec.Cmd
	conn     *grpc.ClientConn
	client   *clients.SourceClient
	errors   int
	warnings int
}

type DestinationPlugin struct {
	cmd    *exec.Cmd
	conn   *grpc.ClientConn
	client *clients.DestinationClient
}

type PluginManager struct {
	logger         zerolog.Logger
	directory      string
	versionsClient *versions.Client
}

type PluginManagerOption func(*PluginManager)

func (p *DestinationPlugin) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	if p.cmd != nil && p.cmd.Process != nil {
		if err := p.cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return nil
}

func (p *DestinationPlugin) GetClient() *clients.DestinationClient {
	return p.client
}

func (p *SourcePlugin) Errors() int {
	return p.errors
}

func (p *SourcePlugin) Warnings() int {
	return p.warnings
}

func (p *SourcePlugin) Close() error {
	if p.conn != nil {
		p.conn.Close()
	}
	if p.cmd != nil && p.cmd.Process != nil {
		if err := p.cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return nil
}

func (p *SourcePlugin) GetClient() *clients.SourceClient {
	return p.client
}

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
		logger:         log.Logger,
		directory:      "./.cq",
		versionsClient: versions.NewClient(),
	}
	// initialize all plugins registry
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// DownloadSource downloads a plugin from the specified registry and return the path to plugin
func (p *PluginManager) DownloadSource(ctx context.Context, spec *specs.Source) (string, error) {
	switch spec.Registry {
	case specs.RegistryLocal, specs.RegistryGrpc:
		fmt.Printf("Skipping plugin download. registry: %s, path: %s\n", spec.Registry, spec.Path)
		p.logger.Info().Str("registry", spec.Registry.String()).Msg("Skiping plugin download")
		return "", nil
	case specs.RegistryGithub:
		return p.downloadSourceGitHub(ctx, spec)
	default:
		return "", fmt.Errorf("unknown registry: %s", spec.Registry)
	}
}

func (p *PluginManager) downloadSourceGitHub(ctx context.Context, spec *specs.Source) (string, error) {
	var err error
	pathSplit := strings.Split(spec.Path, "/")
	org, repo := pathSplit[0], pathSplit[1]
	if spec.Version == "latest" || spec.Version == "" {
		// if version is latest, we need to get the version number from github
		spec.Version, err = p.versionsClient.GetLatestPluginRelease(ctx, org, "source", repo)
		if err != nil {
			return "", err
		}
	}

	pluginName := fmt.Sprintf("cq-source-%s_%s_%s", repo, runtime.GOOS, runtime.GOARCH)
	dirPath := filepath.Join(p.directory, "plugins", spec.Registry.String(), org, repo, spec.Version)
	pluginPath := filepath.Join(dirPath, pluginName)
	if _, err := os.Stat(pluginPath); err == nil {
		fmt.Printf("Plugin already exists at %s. Skipping download.\n", pluginPath)
		p.logger.Info().Str("path", pluginPath).Msg("Plugin already exists. Skipping download.")
		return pluginPath, nil
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create plugin directory %s: %w", dirPath, err)
	}
	// we use convention over configuration and we use github as our registry. Similar to how terraform and homebrew work.
	// For example:
	// https://github.com/cloudquery/cloudquery/releases/download/plugins-source-test-v1.1.0/test_darwin_amd64.zip
	pluginUrl := fmt.Sprintf("https://github.com/cloudquery/cloudquery/releases/download/plugins/source/%s/%s/%s_%s_%s.zip", repo, spec.Version, repo, runtime.GOOS, runtime.GOARCH)
	archivePath := "plugins/source/" + repo
	if org != "cloudquery" {
		// https://github.com/yevgenypats/cq-source-test/releases/download/v1.0.0/cq-source-test_linux_amd64.zip
		pluginUrl = fmt.Sprintf("https://github.com/%s/cq-source-%s/releases/download/%s/cq-source-%s_%s_%s.zip", org, repo, spec.Version, repo, runtime.GOOS, runtime.GOARCH)
		archivePath = "cq-source-" + repo
	}

	fmt.Printf("Downloading plugin from: %s to: %s.zip \n", pluginUrl, pluginPath)
	if err := downloadFile(pluginPath+".zip", pluginUrl); err != nil {
		return "", fmt.Errorf("failed to download plugin: %w", err)
	}
	archive, err := zip.OpenReader(pluginPath + ".zip")
	if err != nil {
		return "", fmt.Errorf("failed to open plugin archive: %w", err)
	}
	fileInArchive, err := archive.Open(archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to open plugin archive plugins/source/%s: %w", repo, err)
	}
	out, err := os.OpenFile(pluginPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", pluginPath, err)
	}
	_, err = io.Copy(out, fileInArchive)
	if err != nil {
		return "", fmt.Errorf("failed to copy body to file: %w", err)
	}
	err = out.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close file: %w", err)
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

func (p *PluginManager) NewSourcePlugin(ctx context.Context, spec *specs.Source) (*SourcePlugin, error) {
	pl := SourcePlugin{}
	// var grpcTarget string
	var pluginPath string
	switch spec.Registry {
	case specs.RegistryGrpc:
		// This is a special case as we dont spawn any process
		conn, err := grpc.Dial(spec.Path, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("failed to dial grpc target %s: %w", spec.Path, err)
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
	cmd := exec.CommandContext(ctx, pluginPath, "serve", "--network", "unix", "--address", grpcTarget,
		"--log-level", p.logger.GetLevel().String(), "--log-format", "json")
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start plugin %s: %w", pluginPath, err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("plugin %s exited with error: %v\n", spec.Path, err)
			p.logger.Error().Err(err).Str("plugin", spec.Path).Msg("plugin exited")
		}
	}()
	pl.cmd = cmd

	go func() {
		scanner := bufio.NewScanner(reader)
		fmt.Println("starting reading")
		for scanner.Scan() {
			fmt.Println("read line")
			var structuredLogLine map[string]interface{}
			b := scanner.Bytes()
			if err := json.Unmarshal(b, &structuredLogLine); err != nil {
				p.logger.Err(err).Str("line", string(b)).Msg("failed to unmarshal log line from plugin")
			} else {
				fmt.Printf("%v", structuredLogLine)
				jsonToLog(&pl, structuredLogLine, p.logger)
			}
		}
	}()
	// remove the socket file if it exists
	// os.Remove(grpcTarget)
	conn, err := grpc.DialContext(ctx, "unix://"+grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("failed to kill plugin", err)
		}
		return &pl, err
	}
	pl.conn = conn
	pl.client = clients.NewSourceClient(conn)
	return &pl, nil
}
