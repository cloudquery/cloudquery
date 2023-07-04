package appengine

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
)

func createApps(gsrv *grpc.Server) error {
	fakeServer := &fakeAppsServer{}
	pb.RegisterApplicationsServer(gsrv, fakeServer)
	return nil
}

type fakeAppsServer struct {
	pb.UnimplementedApplicationsServer
}

func (*fakeAppsServer) GetApplication(context.Context, *pb.GetApplicationRequest) (*pb.Application, error) {
	resp := pb.Application{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestApps(t *testing.T) {
	client.MockTestGrpcHelper(t, Apps(), createApps, client.TestOptions{})
}
