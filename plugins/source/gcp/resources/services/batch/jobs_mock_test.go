package batch

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createJobs(gsrv *grpc.Server) error {
	fakeServer := &fakeJobsServer{}
	pb.RegisterBatchServiceServer(gsrv, fakeServer)
	return nil
}

type fakeJobsServer struct {
	pb.UnimplementedBatchServiceServer
}

func (*fakeJobsServer) ListJobs(context.Context, *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	resp := pb.ListJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestJobs(t *testing.T) {
	client.MockTestGrpcHelper(t, Jobs(), createJobs, client.TestOptions{})
}
