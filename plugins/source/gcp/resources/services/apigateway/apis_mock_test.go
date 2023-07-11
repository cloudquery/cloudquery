package apigateway

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createApis(gsrv *grpc.Server) error {
	fakeServer := &fakeApisServer{}
	pb.RegisterApiGatewayServiceServer(gsrv, fakeServer)
	return nil
}

type fakeApisServer struct {
	pb.UnimplementedApiGatewayServiceServer
}

func (*fakeApisServer) ListApis(context.Context, *pb.ListApisRequest) (*pb.ListApisResponse, error) {
	resp := pb.ListApisResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestApis(t *testing.T) {
	client.MockTestGrpcHelper(t, Apis(), createApis, client.TestOptions{})
}
