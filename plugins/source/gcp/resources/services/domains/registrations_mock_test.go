package domains

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/domains/apiv1beta1/domainspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createRegistrations(gsrv *grpc.Server) error {
	fakeServer := &fakeRegistrationsServer{}
	pb.RegisterDomainsServer(gsrv, fakeServer)
	return nil
}

type fakeRegistrationsServer struct {
	pb.UnimplementedDomainsServer
}

func (*fakeRegistrationsServer) ListRegistrations(context.Context, *pb.ListRegistrationsRequest) (*pb.ListRegistrationsResponse, error) {
	resp := pb.ListRegistrationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestRegistrations(t *testing.T) {
	client.MockTestGrpcHelper(t, Registrations(), createRegistrations, client.TestOptions{})
}
