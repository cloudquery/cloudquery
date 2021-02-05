package sdk

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

func NewLogger(verbose bool, options ...zap.Option) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	disableCaller := true
	if verbose {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
		disableCaller = false
	}
	return zap.Config{
		Sampling:         nil,
		Level:            level,
		Development:      true,
		DisableCaller:    disableCaller,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(options...)
}

func GetProviderPluginClient(path string) (CQProvider, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: PluginMap,
		},
		Cmd:             exec.Command(path),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		SyncStderr: os.Stderr,
		SyncStdout: os.Stdout,
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  hclog.Info,
			Name:   "plugin",
		}),
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

	p := raw.(CQProvider)

	return p, nil
}
