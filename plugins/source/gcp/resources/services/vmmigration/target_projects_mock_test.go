package vmmigration

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
)

func createTargetProjectsServer(gsrv *grpc.Server) error {
	fakeServer := &fakeTargetProjectsServer{}
	pb.RegisterVmMigrationServer(gsrv, fakeServer)
	return nil
}

type fakeTargetProjectsServer struct {
	pb.UnimplementedVmMigrationServer
}

func (*fakeTargetProjectsServer) ListTargetProjects(context.Context, *pb.ListTargetProjectsRequest) (*pb.ListTargetProjectsResponse, error) {
	resp := pb.ListTargetProjectsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestTargetProjects(t *testing.T) {
	client.MockTestGrpcHelper(t, TargetProjects(), createTargetProjectsServer, client.TestOptions{})
}
