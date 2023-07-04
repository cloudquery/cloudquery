package workflows

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
)

func createServer(gsrv *grpc.Server) error {
	fakeServer := &fakeServer{}
	pb.RegisterWorkflowsServer(gsrv, fakeServer)
	return nil
}

type fakeServer struct {
	pb.UnimplementedWorkflowsServer
}

func (*fakeServer) ListWorkflows(context.Context, *pb.ListWorkflowsRequest) (*pb.ListWorkflowsResponse, error) {
	resp := pb.ListWorkflowsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Workflows(), createServer, client.TestOptions{})
}
