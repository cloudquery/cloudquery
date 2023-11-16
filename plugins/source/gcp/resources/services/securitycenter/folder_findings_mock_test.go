package securitycenter

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createFolderFindings(gsrv *grpc.Server) error {
	fakeServer := &fakeFolderSecretsServer{}
	pb.RegisterSecurityCenterServer(gsrv, fakeServer)
	return nil
}

type fakeFolderSecretsServer struct {
	pb.UnimplementedSecurityCenterServer
}

func (*fakeFolderSecretsServer) ListFindings(context.Context, *pb.ListFindingsRequest) (*pb.ListFindingsResponse, error) {
	resp := pb.ListFindingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	// We can't mock this field
	resp.ListFindingsResults[0].Finding.SourceProperties = nil
	return &resp, nil
}

func TestFolderFindings(t *testing.T) {
	client.MockTestGrpcHelper(t, FolderFindings(), createFolderFindings, client.TestOptions{})
}
