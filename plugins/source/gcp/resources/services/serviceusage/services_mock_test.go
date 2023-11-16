package serviceusage

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
)

func createServices(gsrv *grpc.Server) error {
	fakeServer := &fakeServicesServer{}
	pb.RegisterServiceUsageServer(gsrv, fakeServer)
	return nil
}

type fakeServicesServer struct {
	pb.UnimplementedServiceUsageServer
}

func (*fakeServicesServer) ListServices(context.Context, *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	resp := pb.ListServicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestServices(t *testing.T) {
	client.MockTestGrpcHelper(t, Services(), createServices, client.TestOptions{})
}
