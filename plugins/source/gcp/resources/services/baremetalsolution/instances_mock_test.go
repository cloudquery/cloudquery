package baremetalsolution

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createInstances(gsrv *grpc.Server) error {
	fakeServer := &fakeInstancesServer{}
	pb.RegisterBareMetalSolutionServer(gsrv, fakeServer)
	return nil
}

type fakeInstancesServer struct {
	pb.UnimplementedBareMetalSolutionServer
}

func (*fakeInstancesServer) ListInstances(context.Context, *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	resp := pb.ListInstancesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Instances(), createInstances, client.TestOptions{})
}
