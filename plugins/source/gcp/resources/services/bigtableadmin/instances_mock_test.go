package bigtableadmin

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func createInstances(gsrv *grpc.Server) error {
	fakeServer := &fakeInstancesServer{}
	pb.RegisterBigtableInstanceAdminServer(gsrv, fakeServer)
	return nil
}

type fakeInstancesServer struct {
	pb.UnimplementedBigtableInstanceAdminServer
}

func (f *fakeInstancesServer) ListInstances(context.Context, *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	resp := pb.ListInstancesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.Instances[0].Name = "projects/testProject/instances/test-instance"
	resp.FailedLocations = nil
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Instances(), createInstances, client.TestOptions{})
}
