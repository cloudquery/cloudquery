package clouderrorreporting

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
)

func createErrorGroupStats(gsrv *grpc.Server) error {
	fakeServer := &fakeErrorGroupStatsServer{}
	pb.RegisterErrorStatsServiceServer(gsrv, fakeServer)
	return nil
}

type fakeErrorGroupStatsServer struct {
	pb.UnimplementedErrorStatsServiceServer
}

func (*fakeErrorGroupStatsServer) ListGroupStats(context.Context, *pb.ListGroupStatsRequest) (*pb.ListGroupStatsResponse, error) {
	resp := pb.ListGroupStatsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeErrorGroupStatsServer) ListEvents(context.Context, *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {
	resp := pb.ListEventsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestErrorGroupStats(t *testing.T) {
	client.MockTestGrpcHelper(t, ErrorGroupStats(), createErrorGroupStats, client.TestOptions{})
}
