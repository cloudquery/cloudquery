package clouddeploy

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createTargets(gsrv *grpc.Server) error {
	fakeServer := &fakeTargetsServer{}
	pb.RegisterCloudDeployServer(gsrv, fakeServer)
	return nil
}

type fakeTargetsServer struct {
	pb.UnimplementedCloudDeployServer
}

func (*fakeTargetsServer) ListTargets(context.Context, *pb.ListTargetsRequest) (*pb.ListTargetsResponse, error) {
	resp := pb.ListTargetsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestTargets(t *testing.T) {
	client.MockTestGrpcHelper(t, Targets(), createTargets, client.TestOptions{})
}
