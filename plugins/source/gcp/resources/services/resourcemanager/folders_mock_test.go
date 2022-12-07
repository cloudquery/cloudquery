package resourcemanager

import (
	"context"
	"fmt"
	"testing"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createFolders(gsrv *grpc.Server) error {
	fakeServer := &fakeFoldersServer{}
	pb.RegisterFoldersServer(gsrv, fakeServer)

	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := resourcemanager.NewFoldersClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	p, err := createProjects()
	if err != nil {
		return nil, fmt.Errorf("createProjects failed: %w", err)
	}

	return &client.Services{
		ResourcemanagerFoldersClient:  svc,
		ResourcemanagerProjectsClient: p.ResourcemanagerProjectsClient,
	}, nil
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
