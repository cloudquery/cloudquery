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

func createServices(gsrv *grpc.Server) error {
	pb.RegisterServicesServer(gsrv, &fakeServicesServer{})
	pb.RegisterVersionsServer(gsrv, &fakeVersionsServer{})
	pb.RegisterInstancesServer(gsrv, &fakeInstancesServer{})
	return nil
}

type fakeServicesServer struct {
	pb.UnimplementedServicesServer
}

func (*fakeServicesServer) ListServices(context.Context, *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	resp := pb.ListServicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

type fakeVersionsServer struct {
	pb.UnimplementedVersionsServer
}

func (*fakeVersionsServer) ListVersions(context.Context, *pb.ListVersionsRequest) (*pb.ListVersionsResponse, error) {
	resp := pb.ListVersionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

type fakeInstancesServer struct {
	pb.UnimplementedInstancesServer
}

func (*fakeInstancesServer) ListInstances(context.Context, *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	resp := pb.ListInstancesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestServices(t *testing.T) {
	client.MockTestGrpcHelper(t, Services(), createServices, client.TestOptions{})
}
