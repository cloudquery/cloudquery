package plugins

import (
	"os/exec"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
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
		if err := p.cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return nil
}

func (p *DestinationPlugin) GetClient() *clients.DestinationClient {
	return p.client
}

func destJsonToLog(msg map[string]interface{}, l zerolog.Logger) {
	switch msg["level"] {
	case "trace":
		l.Trace().Fields(msg).Msg("")
	case "debug":
		l.Debug().Fields(msg).Msg("")
	case "info":
		l.Debug().Fields(msg).Msg("")
	case "warn":
		l.Warn().Fields(msg).Msg("")
	case "error":
		l.Warn().Fields(msg).Msg("")
	default:
		l.Error().Fields(msg).Msg("unknown level")
	}
}