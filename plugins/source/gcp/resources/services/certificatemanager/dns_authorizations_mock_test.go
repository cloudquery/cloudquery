package certificatemanager

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createDnsAuthorizations(gsrv *grpc.Server) error {
	fakeServer := &fakeDnsAuthorizationsServer{}
	pb.RegisterCertificateManagerServer(gsrv, fakeServer)
	return nil
}

type fakeDnsAuthorizationsServer struct {
	pb.UnimplementedCertificateManagerServer
}

func (*fakeDnsAuthorizationsServer) ListDnsAuthorizations(context.Context, *pb.ListDnsAuthorizationsRequest) (*pb.ListDnsAuthorizationsResponse, error) {
	resp := pb.ListDnsAuthorizationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestDnsAuthorizations(t *testing.T) {
	client.MockTestGrpcHelper(t, DnsAuthorizations(), createDnsAuthorizations, client.TestOptions{})
}
