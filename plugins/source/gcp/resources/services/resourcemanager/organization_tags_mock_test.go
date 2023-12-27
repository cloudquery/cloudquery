package resourcemanager

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createOrganizationTagKeys(gsrv *grpc.Server) error {
	fakeServer := &fakeOrganizationTagKeysServer{}
	pb.RegisterTagKeysServer(gsrv, fakeServer)
	return createOrganizationTagValues(gsrv)
}

type fakeOrganizationTagKeysServer struct {
	pb.UnimplementedTagKeysServer
}

func (*fakeOrganizationTagKeysServer) ListTagKeys(_ context.Context, req *pb.ListTagKeysRequest) (*pb.ListTagKeysResponse, error) {
	resp := pb.ListTagKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}

	for _, f := range resp.TagKeys {
		f.Name = "tagKeys/456"
		f.Parent = req.Parent
	}

	resp.NextPageToken = ""
	return &resp, nil
}

func TestOrganizationTagKeys(t *testing.T) {
	client.MockTestHelper(t, OrganizationTagKeys(), client.WithCreateGrpcService(createOrganizationTagKeys))
}

func createOrganizationTagValues(gsrv *grpc.Server) error {
	fakeServer := &fakeOrganizationTagValuesServer{}
	pb.RegisterTagValuesServer(gsrv, fakeServer)
	return nil
}

type fakeOrganizationTagValuesServer struct {
	pb.UnimplementedTagValuesServer
}

func (*fakeOrganizationTagValuesServer) ListTagValues(_ context.Context, req *pb.ListTagValuesRequest) (*pb.ListTagValuesResponse, error) {
	resp := pb.ListTagValuesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}

	for _, f := range resp.TagValues {
		f.Name = "tagValues/789"
		f.Parent = req.Parent
	}

	resp.NextPageToken = ""
	return &resp, nil
}

func TestOrganizationTagValues(t *testing.T) {
	client.MockTestHelper(t, organizationTagValues(), client.WithCreateGrpcService(createOrganizationTagValues))
}
