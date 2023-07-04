package resourcemanager

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)

func createFolders(gsrv *grpc.Server) error {
	fakeServer := &fakeFoldersServer{}
	pb.RegisterFoldersServer(gsrv, fakeServer)
	return createProjects(gsrv)
}

type fakeFoldersServer struct {
	pb.UnimplementedFoldersServer
}

func (*fakeFoldersServer) ListFolders(context.Context, *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
	resp := pb.ListFoldersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	for _, f := range resp.Folders {
		f.Parent = "organizations/123"
	}

	resp.NextPageToken = ""
	return &resp, nil
}

func TestFolders(t *testing.T) {
	client.MockTestGrpcHelper(t, Folders(), createFolders, client.TestOptions{})
}
