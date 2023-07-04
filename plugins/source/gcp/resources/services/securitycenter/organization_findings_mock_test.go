package securitycenter

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createOrganizationFindings(gsrv *grpc.Server) error {
	fakeServer := &fakeOrganizationSecretsServer{}
	pb.RegisterSecurityCenterServer(gsrv, fakeServer)
	return nil
}

type fakeOrganizationSecretsServer struct {
	pb.UnimplementedSecurityCenterServer
}

func (*fakeOrganizationSecretsServer) ListFindings(context.Context, *pb.ListFindingsRequest) (*pb.ListFindingsResponse, error) {
	resp := pb.ListFindingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	// We can't mock this field
	resp.ListFindingsResults[0].Finding.SourceProperties = nil
	return &resp, nil
}

func TestOrganizationFindings(t *testing.T) {
	client.MockTestGrpcHelper(t, OrganizationFindings(), createOrganizationFindings, client.TestOptions{})
}
