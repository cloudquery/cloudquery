package iam

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"
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

func TestRoles(t *testing.T) {
	client.MockTestGrpcHelper(t, Roles(), createRoles, client.TestOptions{})
}
