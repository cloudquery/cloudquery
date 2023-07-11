package functions

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/functions/apiv1/functionspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createFunctions(gsrv *grpc.Server) error {
	fakeServer := &fakeFunctionsServer{}
	pb.RegisterCloudFunctionsServiceServer(gsrv, fakeServer)
	return nil
}

type fakeFunctionsServer struct {
	pb.UnimplementedCloudFunctionsServiceServer
}

func (*fakeFunctionsServer) ListFunctions(context.Context, *pb.ListFunctionsRequest) (*pb.ListFunctionsResponse, error) {
	resp := pb.ListFunctionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestFunctions(t *testing.T) {
	client.MockTestGrpcHelper(t, Functions(), createFunctions, client.TestOptions{})
}
