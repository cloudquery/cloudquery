package manageddestination

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cloudquery/cloudquery/cli/internal/download"
	"github.com/cloudquery/cloudquery/cli/internal/logging"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultDownloadDir = ".cq"
	maxMsgSize         = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	directory      string
	cmd            *exec.Cmd
	logger         zerolog.Logger
	grpcSocketName string
	wg             *sync.WaitGroup
	Conn           *grpc.ClientConn
	Spec           specs.Destination
	noSentry       bool
}

type Clients []*Client

type Option func(*Client)

func WithLogger(logger zerolog.Logger) func(*Client) {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithDirectory(directory string) func(*Client) {
	return func(c *Client) {
		c.directory = directory
	}
}

func WithNoSentry() func(*Client) {
	return func(c *Client) {
		c.noSentry = true
	}
}

func NewClients(ctx context.Context, destinationSpecs []*specs.Destination, opts ...Option) (Clients, error) {
	clients := make(Clients, len(destinationSpecs))
	for i, spec := range destinationSpecs {
		client, err := NewClient(ctx, *spec, opts...)
		if err != nil {
			return nil, err
		}
		clients[i] = client
	}
	return clients, nil
}

func (c Clients) ClientByName(name string) *Client {
	for _, client := range c {
		if client.Spec.Name == name {
			return client
		}
	}
	return nil
}

func (c Clients) ClientsByNames(names []string) []*Client {
	clients := make([]*Client, len(names))
	for i, client := range c {
		if slices.Contains(names, client.Spec.Name) {
			clients[i] = client
		}
	}
	return clients
}

func (c Clients) Specs() []specs.Destination {
	specs := make([]specs.Destination, len(c))
	for i, client := range c {
		specs[i] = client.Spec
	}
	return specs
}

func (c Clients) Names() []string {
	names := make([]string, len(c))
	for i, client := range c {
		names[i] = client.Spec.Name
	}
	return names
}

func (c Clients) Terminate() error {
	for _, client := range c {
		if err := client.Terminate(); err != nil {
			return err
		}
	}
	return nil
}

// NewClient creates a new plugin client.
// If registrySpec is GitHub then client downloads the plugin, spawns it and creates a gRPC connection.
// If registrySpec is Local then client spawns the plugin and creates a gRPC connection.
// If registrySpec is gRPC then clients creates a new connection
func NewClient(ctx context.Context, spec specs.Destination, opts ...Option) (*Client, error) {
	c := Client{
		directory: defaultDownloadDir,
		wg:        &sync.WaitGroup{},
		Spec:      spec,
	}
	for _, opt := range opts {
		opt(&c)
	}
	var err error
	switch spec.Registry {
	case specs.RegistryGrpc:
		c.Conn, err = grpc.DialContext(ctx, spec.Path,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to dial grpc source plugin at %s: %w", spec.Path, err)
		}
	case specs.RegistryLocal:
		if err := c.startLocal(ctx, spec.Path); err != nil {
			return nil, err
		}
	case specs.RegistryGithub:
		pathSplit := strings.Split(spec.Path, "/")
		if len(pathSplit) != 2 {
			return nil, fmt.Errorf("invalid github plugin path: %s. format should be owner/repo", spec.Path)
		}
		org, name := pathSplit[0], pathSplit[1]
		localPath := filepath.Join(c.directory, "plugins", string(download.PluginTypeDestination), org, name, spec.Version, "plugin")
		localPath = download.WithBinarySuffix(localPath)
		if err := download.DownloadPluginFromGithub(ctx, localPath, org, name, spec.Version, download.PluginTypeDestination); err != nil {
			return nil, err
		}
		if err := c.startLocal(ctx, localPath); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (c *Client) startLocal(ctx context.Context, path string) error {
	c.grpcSocketName = GenerateRandomUnixSocketName()
	// spawn the plugin first and then connect
	args := []string{"serve", "--network", "unix", "--address", c.grpcSocketName,
		"--log-level", c.logger.GetLevel().String(), "--log-format", "json"}
	if c.noSentry {
		args = append(args, "--no-sentry")
	}
	cmd := exec.CommandContext(ctx, path, args...)
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = getSysProcAttr()
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start destination plugin %s: %w", path, err)
	}

	c.cmd = cmd

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		lr := logging.NewLogReader(reader)
		for {
			line, err := lr.NextLine()
			if errors.Is(err, io.EOF) {
				break
			}
			if errors.Is(err, logging.ErrLogLineToLong) {
				c.logger.Info().Str("line", string(line)).Msg("truncated destination plugin log line")
				continue
			}
			if err != nil {
				c.logger.Err(err).Msg("failed to read log line from destination plugin")
				break
			}
			var structuredLogLine map[string]any
			if err := json.Unmarshal(line, &structuredLogLine); err != nil {
				c.logger.Err(err).Str("line", string(line)).Msg("failed to unmarshal log line from destination plugin")
			} else {
				logging.JSONToLog(c.logger, structuredLogLine)
			}
		}
	}()

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		d := &net.Dialer{}
		return d.DialContext(ctx, "unix", addr)
	}
	c.Conn, err = grpc.DialContext(ctx, c.grpcSocketName, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithContextDialer(dialer))
	if err != nil {
		if err := cmd.Process.Kill(); err != nil {
			c.logger.Error().Err(err).Msg("failed to kill plugin process")
		}
		return err
	}
	return nil
}

func (c *Client) Terminate() error {
	// wait for log streaming to complete before returning from this function
	defer c.wg.Wait()

	if c.grpcSocketName != "" {
		defer func() {
			if err := os.RemoveAll(c.grpcSocketName); err != nil {
				c.logger.Error().Err(err).Msg("failed to remove source socket file")
			}
		}()
	}

	if c.Conn != nil {
		if err := c.Conn.Close(); err != nil {
			c.logger.Error().Err(err).Msg("failed to close gRPC connection to source plugin")
		}
		c.Conn = nil
	}
	if c.cmd != nil && c.cmd.Process != nil {
		if err := c.terminateProcess(); err != nil {
			return err
		}
	}

	return nil
}
