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

func createFeaturestoreLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeFeaturestoreLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeFeaturestoreLocationsRelationsServer{}
	aiplatformpb.RegisterFeaturestoreServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeFeaturestoreLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeFeaturestoreLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestFeaturestoreLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, FeaturestoreLocations(), createFeaturestoreLocations, client.TestOptions{})
}

type fakeFeaturestoreLocationsRelationsServer struct {
	aiplatformpb.UnimplementedFeaturestoreServiceServer
}

func (*fakeFeaturestoreLocationsRelationsServer) ListEntityTypes(context.Context, *aiplatformpb.ListEntityTypesRequest) (*aiplatformpb.ListEntityTypesResponse, error) {
	resp := aiplatformpb.ListEntityTypesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeFeaturestoreLocationsRelationsServer) ListFeatures(context.Context, *aiplatformpb.ListFeaturesRequest) (*aiplatformpb.ListFeaturesResponse, error) {
	resp := aiplatformpb.ListFeaturesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeFeaturestoreLocationsRelationsServer) ListFeaturestores(context.Context, *aiplatformpb.ListFeaturestoresRequest) (*aiplatformpb.ListFeaturestoresResponse, error) {
	resp := aiplatformpb.ListFeaturestoresResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
