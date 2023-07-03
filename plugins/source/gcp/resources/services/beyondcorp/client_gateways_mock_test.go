package beyondcorp

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createClientGateways(gsrv *grpc.Server) error {
	fakeServer := &fakeClientGatewaysServer{}
	pb.RegisterClientGatewaysServiceServer(gsrv, fakeServer)
	return nil
}

type fakeClientGatewaysServer struct {
	pb.UnimplementedClientGatewaysServiceServer
}

func (*fakeClientGatewaysServer) ListClientGateways(context.Context, *pb.ListClientGatewaysRequest) (*pb.ListClientGatewaysResponse, error) {
	resp := pb.ListClientGatewaysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestClientGateways(t *testing.T) {
	client.MockTestGrpcHelper(t, ClientGateways(), createClientGateways, client.TestOptions{})
}
