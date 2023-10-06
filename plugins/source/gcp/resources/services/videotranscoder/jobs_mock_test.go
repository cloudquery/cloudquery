package videotranscoder

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
)

func createJobsServer(gsrv *grpc.Server) error {
	fakeServer := &fakeJobsServer{}
	pb.RegisterTranscoderServiceServer(gsrv, fakeServer)
	return nil
}

type fakeJobsServer struct {
	pb.UnimplementedTranscoderServiceServer
}

func (*fakeJobsServer) ListJobs(context.Context, *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	resp := pb.ListJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Jobs(), createJobsServer, client.TestOptions{})
}
