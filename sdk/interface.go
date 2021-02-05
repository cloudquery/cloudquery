package sdk

import (
	"context"
	"github.com/cloudquery/cloudquery/sdk/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	MagicCookieKey:   "CQ_PLUGIN_COOKIE",
	MagicCookieValue: "6753812e-79c2-4af5-ad01-e6083c374e1f",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"provider": &CQPlugin{},
}

type CQProvider interface {
	Init(driver string, dsn string, verbose bool) error
	Fetch(data []byte) error
	GenConfig() (string, error)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type CQPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl CQProvider
}

func (p *CQPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterProviderServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *CQPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewProviderClient(c)}, nil
}

func ServePlugin(provider CQProvider) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		VersionedPlugins: map[int]plugin.PluginSet {
		1: {
			"provider": &CQPlugin{Impl: provider},
		}},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}