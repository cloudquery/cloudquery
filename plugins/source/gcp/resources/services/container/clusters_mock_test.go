package container

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"cloud.google.com/go/container/apiv1"

	pb "google.golang.org/genproto/googleapis/container/v1"

	"google.golang.org/api/option"
)

func createClusters() (*client.Services, error) {
	fakeServer := &fakeClustersServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterClusterManagerServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := container.NewClusterManagerClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		ContainerClusterManagerClient: svc,
	}, nil
}

type fakeClustersServer struct {
	pb.UnimplementedClusterManagerServer
}

func (f *fakeClustersServer) ListClusters(context.Context, *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	resp := pb.ListClustersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestClusters(t *testing.T) {
	client.MockTestHelper(t, Clusters(), createClusters, client.TestOptions{})
}
