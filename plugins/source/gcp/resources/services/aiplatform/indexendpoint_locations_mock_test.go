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

func createIndexendpointLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeIndexendpointLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeIndexendpointLocationsRelationsServer{}
	aiplatformpb.RegisterIndexEndpointServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeIndexendpointLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeIndexendpointLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestIndexendpointLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, IndexendpointLocations(), createIndexendpointLocations, client.TestOptions{})
}

type fakeIndexendpointLocationsRelationsServer struct {
	aiplatformpb.UnimplementedIndexEndpointServiceServer
}

func (*fakeIndexendpointLocationsRelationsServer) ListIndexEndpoints(context.Context, *aiplatformpb.ListIndexEndpointsRequest) (*aiplatformpb.ListIndexEndpointsResponse, error) {
	resp := aiplatformpb.ListIndexEndpointsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
