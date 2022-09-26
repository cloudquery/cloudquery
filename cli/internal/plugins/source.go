package plugins

import (
	"fmt"
	"os/exec"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type SourcePlugin struct {
	cmd      *exec.Cmd
	conn     *grpc.ClientConn
	client   *clients.SourceClient
	errors   int
	warnings int
}

func (p *SourcePlugin) jsonToLog(msg map[string]interface{}, l zerolog.Logger) {
	switch msg["level"] {
	case "trace":
		l.Trace().Fields(msg).Msg("")
	case "debug":
		l.Debug().Fields(msg).Msg("")
	case "info":
		l.Debug().Fields(msg).Msg("")
	case "warn":
		p.warnings++
		l.Warn().Fields(msg).Msg("")
	case "error":
		p.errors++
		l.Warn().Fields(msg).Msg("")
	default:
		l.Error().Fields(msg).Msg("unknown level")
	}
}

func (p *SourcePlugin) Errors() int {
	return p.errors
}

func (p *SourcePlugin) Warnings() int {
	return p.warnings
}

func (p *SourcePlugin) Close() error {
	if p.conn != nil {
		if err := p.conn.Close(); err != nil {
			fmt.Println(err)
		}
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
