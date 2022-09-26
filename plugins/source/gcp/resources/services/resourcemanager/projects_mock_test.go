package resourcemanager

import (
	"context"
	"fmt"
	"net"
	"testing"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createProjects() (*client.Services, error) {
	fakeServer := &fakeProjectsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterProjectsServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := resourcemanager.NewProjectsClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		ResourcemanagerProjectsClient: svc,
	}, nil
}

type fakeProjectsServer struct {
	pb.UnimplementedProjectsServer
}

func (*fakeProjectsServer) GetProject(context.Context, *pb.GetProjectRequest) (*pb.Project, error) {
	resp := pb.Project{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), createProjects, client.TestOptions{})
}
