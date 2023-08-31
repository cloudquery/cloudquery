package resourcemanager

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createProjects(gsrv *grpc.Server) error {
	fakeServer := &fakeProjectsServer{}
	resourcemanagerpb.RegisterProjectsServer(gsrv, fakeServer)
	return nil
}

type fakeProjectsServer struct {
	resourcemanagerpb.UnimplementedProjectsServer
}

func (*fakeProjectsServer) GetProject(context.Context, *resourcemanagerpb.GetProjectRequest) (*resourcemanagerpb.Project, error) {
	resp := resourcemanagerpb.Project{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestProjects(t *testing.T) {
	client.MockTestGrpcHelper(t, Projects(), createProjects, client.TestOptions{})
}
