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

func createServiceAccounts(gsrv *grpc.Server) error {
	fakeServer := &fakeServiceAccountsServer{}
	pb.RegisterIAMServer(gsrv, fakeServer)
	return nil
}

type fakeServiceAccountsServer struct {
	pb.UnimplementedIAMServer
}

func (*fakeServiceAccountsServer) ListServiceAccounts(context.Context, *pb.ListServiceAccountsRequest) (*pb.ListServiceAccountsResponse, error) {
	resp := pb.ListServiceAccountsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServiceAccountsServer) ListServiceAccountKeys(context.Context, *pb.ListServiceAccountKeysRequest) (*pb.ListServiceAccountKeysResponse, error) {
	resp := pb.ListServiceAccountKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestServiceAccounts(t *testing.T) {
	client.MockTestGrpcHelper(t, ServiceAccounts(), createServiceAccounts, client.TestOptions{})
}
