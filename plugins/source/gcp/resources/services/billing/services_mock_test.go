package billing

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createServices(gsrv *grpc.Server) error {
	fakeServer := &fakeServicesServer{}
	pb.RegisterCloudCatalogServer(gsrv, fakeServer)
	return nil
}

type fakeServicesServer struct {
	pb.UnimplementedCloudCatalogServer
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
