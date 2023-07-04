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

func createNetworks(gsrv *grpc.Server) error {
	fakeServer := &fakeNetworksServer{}
	pb.RegisterBareMetalSolutionServer(gsrv, fakeServer)
	return nil
}

type fakeNetworksServer struct {
	pb.UnimplementedBareMetalSolutionServer
}

func (*fakeNetworksServer) ListNetworks(context.Context, *pb.ListNetworksRequest) (*pb.ListNetworksResponse, error) {
	resp := pb.ListNetworksResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestNetworks(t *testing.T) {
	client.MockTestGrpcHelper(t, Networks(), createNetworks, client.TestOptions{})
}
