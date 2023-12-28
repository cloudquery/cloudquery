package resourcemanager

import (
	"fmt"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func createProjectTagKeys(gsrv *grpc.Server) error {
	fakeServer := &fakeProjectTagKeysServer{}
	pb.RegisterTagKeysServer(gsrv, fakeServer)
	return createProjectTagValues(gsrv)
}

type fakeProjectTagKeysServer struct {
	pb.UnimplementedTagKeysServer
}

func (*fakeProjectTagKeysServer) ListTagKeys(_ context.Context, req *pb.ListTagKeysRequest) (*pb.ListTagKeysResponse, error) {
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

func TestProjectTagKeys(t *testing.T) {
	client.MockTestHelper(t, ProjectTagKeys(), client.WithCreateGrpcService(createProjectTagKeys))
}

func createProjectTagValues(gsrv *grpc.Server) error {
	fakeServer := &fakeProjectTagValuesServer{}
	pb.RegisterTagValuesServer(gsrv, fakeServer)
	return nil
}

type fakeProjectTagValuesServer struct {
	pb.UnimplementedTagValuesServer
}

func (*fakeProjectTagValuesServer) ListTagValues(_ context.Context, req *pb.ListTagValuesRequest) (*pb.ListTagValuesResponse, error) {
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

func TestProjectTagValues(t *testing.T) {
	client.MockTestHelper(t, projectTagValues(), client.WithCreateGrpcService(createProjectTagValues))
}
