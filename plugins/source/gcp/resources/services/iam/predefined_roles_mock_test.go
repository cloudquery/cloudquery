package iam

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createRoles(gsrv *grpc.Server) error {
	fakeServer := &fakeRolesServer{}
	pb.RegisterIAMServer(gsrv, fakeServer)
	return nil
}

type fakeRolesServer struct {
	pb.UnimplementedIAMServer
}

func (*fakeRolesServer) ListRoles(context.Context, *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	resp := pb.ListRolesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
func TestPredefinedRoles(t *testing.T) {
	client.MockTestHelper(t, PredefinedRoles(), client.WithCreateGrpcService(createRoles))
}
