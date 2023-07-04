package vmmigration

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
)

func createGroupsServer(gsrv *grpc.Server) error {
	fakeServer := &fakeGroupsServer{}
	pb.RegisterVmMigrationServer(gsrv, fakeServer)
	return nil
}

type fakeGroupsServer struct {
	pb.UnimplementedVmMigrationServer
}

func (*fakeGroupsServer) ListGroups(context.Context, *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	resp := pb.ListGroupsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestGroups(t *testing.T) {
	client.MockTestGrpcHelper(t, Groups(), createGroupsServer, client.TestOptions{})
}
