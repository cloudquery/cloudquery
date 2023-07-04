package iam

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/iam/apiv2/iampb"
)

func createDenyPolicies(gsrv *grpc.Server) error {
	fakeServer := &fakeDenyPoliciesServer{}
	pb.RegisterPoliciesServer(gsrv, fakeServer)
	return nil
}

type fakeDenyPoliciesServer struct {
	pb.UnimplementedPoliciesServer
}

func (*fakeDenyPoliciesServer) ListPolicies(context.Context, *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	resp := pb.ListPoliciesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestDenyPolicies(t *testing.T) {
	client.MockTestGrpcHelper(t, DenyPolicies(), createDenyPolicies, client.TestOptions{})
}
