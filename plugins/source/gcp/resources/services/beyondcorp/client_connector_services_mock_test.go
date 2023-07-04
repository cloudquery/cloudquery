package beyondcorp

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createClientConnectorServices(gsrv *grpc.Server) error {
	fakeServer := &fakeClientConnectorServicesServer{}
	pb.RegisterClientConnectorServicesServiceServer(gsrv, fakeServer)
	return nil
}

type fakeClientConnectorServicesServer struct {
	pb.UnimplementedClientConnectorServicesServiceServer
}

func (*fakeClientConnectorServicesServer) ListClientConnectorServices(context.Context, *pb.ListClientConnectorServicesRequest) (*pb.ListClientConnectorServicesResponse, error) {
	resp := pb.ListClientConnectorServicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestClientConnectorServices(t *testing.T) {
	client.MockTestGrpcHelper(t, ClientConnectorServices(), createClientConnectorServices, client.TestOptions{})
}
