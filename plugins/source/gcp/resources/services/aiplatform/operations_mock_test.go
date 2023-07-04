package aiplatform

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createOperations(gsrv *grpc.Server) error {
	fakeServer := &fakeOperationsServer{}
	pb.RegisterOperationsServer(gsrv, fakeServer)
	return nil
}

type fakeOperationsServer struct {
	pb.UnimplementedOperationsServer
}

func (*fakeOperationsServer) ListOperations(context.Context, *pb.ListOperationsRequest) (*pb.ListOperationsResponse, error) {
	resp := pb.ListOperationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestOperations(t *testing.T) {
	client.MockTestGrpcHelper(t, Operations(), createOperations, client.TestOptions{})
}
