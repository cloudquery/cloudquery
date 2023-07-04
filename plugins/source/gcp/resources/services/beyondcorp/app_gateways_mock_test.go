package beyondcorp

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAppGateways(gsrv *grpc.Server) error {
	fakeServer := &fakeAppGatewaysServer{}
	pb.RegisterAppGatewaysServiceServer(gsrv, fakeServer)
	return nil
}

type fakeAppGatewaysServer struct {
	pb.UnimplementedAppGatewaysServiceServer
}

func (*fakeAppGatewaysServer) ListAppGateways(context.Context, *pb.ListAppGatewaysRequest) (*pb.ListAppGatewaysResponse, error) {
	resp := pb.ListAppGatewaysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAppGateways(t *testing.T) {
	client.MockTestGrpcHelper(t, AppGateways(), createAppGateways, client.TestOptions{})
}
