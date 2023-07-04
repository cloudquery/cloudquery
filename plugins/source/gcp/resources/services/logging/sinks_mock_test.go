package logging

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createSinks(gsrv *grpc.Server) error {
	fakeServer := &fakeSinksServer{}
	pb.RegisterConfigServiceV2Server(gsrv, fakeServer)
	return nil
}

type fakeSinksServer struct {
	pb.UnimplementedConfigServiceV2Server
}

func (*fakeSinksServer) ListSinks(context.Context, *pb.ListSinksRequest) (*pb.ListSinksResponse, error) {
	resp := pb.ListSinksResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSinks(t *testing.T) {
	client.MockTestGrpcHelper(t, Sinks(), createSinks, client.TestOptions{})
}
