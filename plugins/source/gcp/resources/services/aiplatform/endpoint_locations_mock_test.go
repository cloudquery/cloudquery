package aiplatform

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
)

func createEndpointLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeEndpointLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeEndpointLocationsRelationsServer{}
	aiplatformpb.RegisterEndpointServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeEndpointLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeEndpointLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestEndpointLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, EndpointLocations(), createEndpointLocations, client.TestOptions{})
}

type fakeEndpointLocationsRelationsServer struct {
	aiplatformpb.UnimplementedEndpointServiceServer
}

func (*fakeEndpointLocationsRelationsServer) ListEndpoints(context.Context, *aiplatformpb.ListEndpointsRequest) (*aiplatformpb.ListEndpointsResponse, error) {
	resp := aiplatformpb.ListEndpointsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
