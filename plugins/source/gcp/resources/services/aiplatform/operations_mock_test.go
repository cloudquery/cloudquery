package aiplatform

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
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
	fakeAny, _ := anypb.New(&longrunningpb.Operation{
		Name: "fake_operation",
	})

	resp.Operations[0].Result = &pb.Operation_Response{Response: fakeAny}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestOperations(t *testing.T) {
	client.MockTestHelper(t, Operations(), client.WithCreateGrpcService(createOperations))
}
