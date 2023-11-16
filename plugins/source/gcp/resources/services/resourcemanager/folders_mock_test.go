package resourcemanager

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
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

func (*fakeFoldersServer) ListFolders(_ context.Context, req *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
	resp := pb.ListFoldersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}

	if strings.HasPrefix(req.Parent, "organizations/") {
		for _, f := range resp.Folders {
			f.Name = "folder/456"
			f.Parent = req.Parent
		}
	} else if req.Parent == "folder/456" {
		for _, f := range resp.Folders {
			f.Name = "folder/789"
			f.Parent = req.Parent
		}
	} else {
		resp.Folders = nil
	}

	resp.NextPageToken = ""
	return &resp, nil
}

func TestFolders(t *testing.T) {
	client.MockTestGrpcHelper(t, Folders(), createFolders, client.TestOptions{})
}
