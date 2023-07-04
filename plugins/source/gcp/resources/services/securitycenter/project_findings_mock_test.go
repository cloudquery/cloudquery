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

func createProjectFindings(gsrv *grpc.Server) error {
	fakeServer := &fakeProjectsSecretsServer{}
	pb.RegisterSecurityCenterServer(gsrv, fakeServer)
	return nil
}

type fakeProjectsSecretsServer struct {
	pb.UnimplementedSecurityCenterServer
}

func (*fakeProjectsSecretsServer) ListFindings(context.Context, *pb.ListFindingsRequest) (*pb.ListFindingsResponse, error) {
	resp := pb.ListFindingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	// We can't mock this field
	resp.ListFindingsResults[0].Finding.SourceProperties = nil
	return &resp, nil
}

func TestProjectFindings(t *testing.T) {
	client.MockTestGrpcHelper(t, OrganizationFindings(), createProjectFindings, client.TestOptions{})
}
