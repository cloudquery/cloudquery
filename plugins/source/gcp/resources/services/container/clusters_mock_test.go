package container

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createClusters(gsrv *grpc.Server) error {
	fakeServer := &fakeClustersServer{}
	pb.RegisterClusterManagerServer(gsrv, fakeServer)
	return nil
}

type fakeClustersServer struct {
	pb.UnimplementedClusterManagerServer
}

func (*fakeClustersServer) ListClusters(context.Context, *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	resp := pb.ListClustersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestClusters(t *testing.T) {
	client.MockTestGrpcHelper(t, Clusters(), createClusters, client.TestOptions{})
}
