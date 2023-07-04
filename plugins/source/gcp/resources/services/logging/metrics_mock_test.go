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

func createMetrics(gsrv *grpc.Server) error {
	fakeServer := &fakeMetricsServer{}
	pb.RegisterMetricsServiceV2Server(gsrv, fakeServer)
	return nil
}

type fakeMetricsServer struct {
	pb.UnimplementedMetricsServiceV2Server
}

func (*fakeMetricsServer) ListLogMetrics(context.Context, *pb.ListLogMetricsRequest) (*pb.ListLogMetricsResponse, error) {
	resp := pb.ListLogMetricsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestMetrics(t *testing.T) {
	client.MockTestGrpcHelper(t, Metrics(), createMetrics, client.TestOptions{})
}
