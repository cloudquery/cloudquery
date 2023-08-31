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

func createVizierLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeVizierLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeVizierLocationsRelationsServer{}
	aiplatformpb.RegisterVizierServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeVizierLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeVizierLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestVizierLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, VizierLocations(), createVizierLocations, client.TestOptions{})
}

type fakeVizierLocationsRelationsServer struct {
	aiplatformpb.UnimplementedVizierServiceServer
}

func (*fakeVizierLocationsRelationsServer) ListStudies(context.Context, *aiplatformpb.ListStudiesRequest) (*aiplatformpb.ListStudiesResponse, error) {
	resp := aiplatformpb.ListStudiesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeVizierLocationsRelationsServer) ListTrials(context.Context, *aiplatformpb.ListTrialsRequest) (*aiplatformpb.ListTrialsResponse, error) {
	resp := aiplatformpb.ListTrialsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
