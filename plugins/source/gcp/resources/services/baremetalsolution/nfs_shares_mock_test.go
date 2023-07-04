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

func createNfsShares(gsrv *grpc.Server) error {
	fakeServer := &fakeNfsSharesServer{}
	pb.RegisterBareMetalSolutionServer(gsrv, fakeServer)
	return nil
}

type fakeNfsSharesServer struct {
	pb.UnimplementedBareMetalSolutionServer
}

func (*fakeNfsSharesServer) ListNfsShares(context.Context, *pb.ListNfsSharesRequest) (*pb.ListNfsSharesResponse, error) {
	resp := pb.ListNfsSharesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestNfsShares(t *testing.T) {
	client.MockTestGrpcHelper(t, NfsShares(), createNfsShares, client.TestOptions{})
}
