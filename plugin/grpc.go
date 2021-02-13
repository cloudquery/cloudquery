package plugin

import (
	"context"
	"github.com/cloudquery/cloudquery/plugin/proto"
)


type GRPCClient struct{ client proto.ProviderClient }

func (m *GRPCClient) Init(driver string, dsn string, verbose bool) error {
	_, err := m.client.Init(context.Background(), &proto.InitRequest{
		Driver: driver,
		Dsn: dsn,
		Verbose: verbose,
	})
	return err
}

func (m *GRPCClient) GenConfig() (string, error) {
	res, err := m.client.GenConfig(context.Background(), &proto.GenConfigRequest{
	})
	if err != nil {
		return "", err
	}
	return res.Yaml, nil
}

func (m *GRPCClient) Fetch(data []byte) error {
	_, err := m.client.Fetch(context.Background(), &proto.FetchRequest{
		Data: data,
	})
	return err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl CQProvider
	proto.UnimplementedProviderServer
}


func (m *GRPCServer) Init(ctx context.Context, req *proto.InitRequest) (*proto.InitResponse, error) {
	return &proto.InitResponse{}, m.Impl.Init(req.Driver, req.Dsn, req.Verbose)
}

func (m *GRPCServer) GenConfig(ctx context.Context, req *proto.GenConfigRequest) (*proto.GenConfigResponse, error) {
	r, err := m.Impl.GenConfig()
	if err != nil {
		return nil, err
	}
	return &proto.GenConfigResponse{Yaml: r}, nil
}

func (m *GRPCServer) Fetch(ctx context.Context, req *proto.FetchRequest) (*proto.FetchResponse, error) {
	err := m.Impl.Fetch(req.Data)
	return &proto.FetchResponse{}, err
}
