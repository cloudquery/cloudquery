package beyondcorp

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAppConnections(gsrv *grpc.Server) error {
	fakeServer := &fakeAppConnectionsServer{}
	pb.RegisterAppConnectionsServiceServer(gsrv, fakeServer)
	return nil
}

type fakeAppConnectionsServer struct {
	pb.UnimplementedAppConnectionsServiceServer
}

func (*fakeAppConnectionsServer) ListAppConnections(context.Context, *pb.ListAppConnectionsRequest) (*pb.ListAppConnectionsResponse, error) {
	resp := pb.ListAppConnectionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAppConnections(t *testing.T) {
	client.MockTestGrpcHelper(t, AppConnections(), createAppConnections, client.TestOptions{})
}
