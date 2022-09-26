// pluginmanager takes care of the lifecycle of plugins,
// including downloading, upgrading, spawning and closing.
// Currently we use GitHub releases as our plugin store.
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
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PluginType string

const (
	PluginTypeSource PluginType = "source"
	PluginTypeDestination PluginType = "destination"
)

type PluginManager struct {
	logger         zerolog.Logger
	directory      string
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
		logger:         log.Logger,
		directory:      "./.cq",
	}
	// initialize all plugins registry
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *PluginManager) NewSourcePlugin(ctx context.Context, registry specs.Registry, path string, version string) (*SourcePlugin, error) {
	pl := SourcePlugin{}
	var pluginPath string
	switch registry {
	case specs.RegistryGrpc:
		// This is a special case as we dont spawn any process
		conn, err := grpc.Dial(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("failed to dial grpc source plugin at %s: %w", path, err)
		}
		pl.conn = conn
		pl.client = clients.NewSourceClient(conn)
		return &pl, nil
	case specs.RegistryLocal:
		pluginPath = path
	case specs.RegistryGithub:
		var err error
		pluginPath, err = p.downloadPluginFromGitHub(ctx, path, version, PluginTypeSource)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown registry: %s", registry)
	}
	grpcTarget := generateRandomUnixSocketName()
	// spawn the plugin first and then connect
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
			fmt.Printf("plugin %s exited with error: %v\n", path, err)
			p.logger.Error().Err(err).Str("plugin", path).Msg("plugin exited")
		}
	}()
	pl.cmd = cmd

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			var structuredLogLine map[string]interface{}
			b := scanner.Bytes()
			if err := json.Unmarshal(b, &structuredLogLine); err != nil {
				p.logger.Err(err).Str("line", string(b)).Msg("failed to unmarshal log line from plugin")
			} else {
				pl.jsonToLog(structuredLogLine, p.logger)
			}
		}
	}()
	
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

func (p *PluginManager) NewDestinationPlugin(ctx context.Context, registry specs.Registry, path string, version string) (*DestinationPlugin, error) {
	pl := DestinationPlugin{}
	var pluginPath string
	switch registry {
	case specs.RegistryGrpc:
		// This is a special case as we dont spawn any process
		conn, err := grpc.Dial(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("failed to dial grpc to destination at %s: %w", path, err)
		}
		pl.conn = conn
		pl.client = clients.NewDestinationClient(conn)
		return &pl, nil
	case specs.RegistryLocal:
		pluginPath = path
	case specs.RegistryGithub:
		var err error
		pluginPath, err = p.downloadPluginFromGitHub(ctx, path, version, PluginTypeDestination)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown registry: %s", registry)
	}
	grpcTarget := generateRandomUnixSocketName()
	// spawn the plugin first and then connect
	cmd := exec.CommandContext(ctx, pluginPath, "serve", "--network", "unix", "--address", grpcTarget,
		"--log-level", p.logger.GetLevel().String(), "--log-format", "json")
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start destination plugin %s: %w", pluginPath, err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("destination plugin %s exited with error: %v\n", path, err)
			p.logger.Error().Err(err).Str("plugin", path).Msg("destination plugin exited")
		}
	}()
	pl.cmd = cmd

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			var structuredLogLine map[string]interface{}
			b := scanner.Bytes()
			if err := json.Unmarshal(b, &structuredLogLine); err != nil {
				p.logger.Err(err).Str("line", string(b)).Msg("failed to unmarshal log line from plugin")
			} else {
				destJsonToLog( structuredLogLine, p.logger)
			}
		}
	}()

	conn, err := grpc.DialContext(ctx, "unix://"+grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("failed to kill destination plugin", err)
		}
		return &pl, err
	}
	pl.conn = conn
	pl.client = clients.NewDestinationClient(conn)
	return &pl, nil
}

func (pm *PluginManager) downloadPluginFromGitHub(ctx context.Context, remotePath string, version string, typ PluginType ) (string, error) {
	var err error
	pathSplit := strings.Split(remotePath, "/")
	org, name := pathSplit[0], pathSplit[1]
	if version == "latest" || version == "" {
		// if version is latest, we need to get the version number from github
		version, err = GetLatestPluginRelease(ctx, org, name, typ)
		if err != nil {
			return "", err
		}
	}
	downloadDir := filepath.Join(pm.directory, "plugins", string(typ), org, name, version)
	pluginPath := path.Join(downloadDir, "plugin")
	pluginZipPath := pluginPath + ".zip"
	// https://github.com/cloudquery/cloudquery/releases/download/plugins-source-test-v1.1.5/test_darwin_amd64.zip
	downloadUrl := fmt.Sprintf("https://github.com/cloudquery/cloudquery/releases/download/plugins-%s-%s-%s/%s_%s_%s.zip", typ, name, version, name, runtime.GOOS, runtime.GOARCH)
	if org != "cloudquery" {
		// https://github.com/yevgenypats/cq-source-test/releases/download/v1.0.1/cq-source-test_darwin_amd64.zip
		downloadUrl = fmt.Sprintf("https://github.com/%s/cq-%s-%s/releases/download/%s/cq-%s-%s_%s_%s.zip", org, typ, name, version, typ, name, runtime.GOOS, runtime.GOARCH)
	}

	if _, err := os.Stat(pluginPath); err == nil {
		pm.logger.Info().Str("path", pluginPath).Msg("Plugin already exists. Skipping download.")
		return pluginPath, nil
	}

	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create plugin directory %s: %w", downloadDir, err)
	}

	err = downloadFile(pluginZipPath, downloadUrl)
	if err != nil {
		return "", fmt.Errorf("failed to download plugin: %w", err)
	}

	archive, err := zip.OpenReader(pluginZipPath)
	if err != nil {
		return "", fmt.Errorf("failed to open plugin archive: %w", err)
	}

	pathInArchive := fmt.Sprintf("plugins/%s/%s", typ, name)
	if org != "cloudquery" {
		pathInArchive = fmt.Sprintf("cq-%s-%s", typ, name)
	}

	fileInArchive, err := archive.Open(pathInArchive)
	if err != nil {
		return "", fmt.Errorf("failed to open plugin archive plugins/source/%s: %w", name, err)
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

