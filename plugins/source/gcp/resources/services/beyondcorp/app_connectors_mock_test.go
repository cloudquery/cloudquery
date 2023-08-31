package beyondcorp

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAppConnectors(gsrv *grpc.Server) error {
	fakeServer := &fakeAppConnectorsServer{}
	pb.RegisterAppConnectorsServiceServer(gsrv, fakeServer)
	return nil
}

type fakeAppConnectorsServer struct {
	pb.UnimplementedAppConnectorsServiceServer
}

func (*fakeAppConnectorsServer) ListAppConnectors(context.Context, *pb.ListAppConnectorsRequest) (*pb.ListAppConnectorsResponse, error) {
	resp := pb.ListAppConnectorsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAppConnectors(t *testing.T) {
	client.MockTestGrpcHelper(t, AppConnectors(), createAppConnectors, client.TestOptions{})
}
