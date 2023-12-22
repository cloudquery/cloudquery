package logging

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
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
	resp.Sinks[0].Options = &pb.LogSink_BigqueryOptions{BigqueryOptions: &pb.BigQueryOptions{UsePartitionedTables: true}}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSinks(t *testing.T) {
	client.MockTestHelper(t, Sinks(), client.WithCreateGrpcService(createSinks))
}
