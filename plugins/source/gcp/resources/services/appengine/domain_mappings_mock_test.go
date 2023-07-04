package appengine

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createDomainMappings(gsrv *grpc.Server) error {
	fakeServer := &fakeDomainMappingsServer{}
	pb.RegisterDomainMappingsServer(gsrv, fakeServer)
	return nil
}

type fakeDomainMappingsServer struct {
	pb.UnimplementedDomainMappingsServer
}

func (*fakeDomainMappingsServer) ListDomainMappings(context.Context, *pb.ListDomainMappingsRequest) (*pb.ListDomainMappingsResponse, error) {
	resp := pb.ListDomainMappingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestDomainMappings(t *testing.T) {
	client.MockTestGrpcHelper(t, DomainMappings(), createDomainMappings, client.TestOptions{})
}
