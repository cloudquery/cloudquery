package videotranscoder

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
)

func createJobTemplatesServer(gsrv *grpc.Server) error {
	fakeServer := &fakeJobTemplatesServer{}
	pb.RegisterTranscoderServiceServer(gsrv, fakeServer)
	return nil
}

type fakeJobTemplatesServer struct {
	pb.UnimplementedTranscoderServiceServer
}

func (*fakeJobTemplatesServer) ListJobTemplates(context.Context, *pb.ListJobTemplatesRequest) (*pb.ListJobTemplatesResponse, error) {
	resp := pb.ListJobTemplatesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestJobTemplates(t *testing.T) {
	client.MockTestGrpcHelper(t, JobTemplates(), createJobTemplatesServer, client.TestOptions{})
}
