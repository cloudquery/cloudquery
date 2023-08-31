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

func createDatasetLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeDatasetLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeDatasetLocationsRelationsServer{}
	aiplatformpb.RegisterDatasetServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeDatasetLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeDatasetLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestDatasetLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, DatasetLocations(), createDatasetLocations, client.TestOptions{})
}

type fakeDatasetLocationsRelationsServer struct {
	aiplatformpb.UnimplementedDatasetServiceServer
}

func (*fakeDatasetLocationsRelationsServer) ListAnnotations(context.Context, *aiplatformpb.ListAnnotationsRequest) (*aiplatformpb.ListAnnotationsResponse, error) {
	resp := aiplatformpb.ListAnnotationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDatasetLocationsRelationsServer) ListDataItems(context.Context, *aiplatformpb.ListDataItemsRequest) (*aiplatformpb.ListDataItemsResponse, error) {
	resp := aiplatformpb.ListDataItemsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDatasetLocationsRelationsServer) ListDatasets(context.Context, *aiplatformpb.ListDatasetsRequest) (*aiplatformpb.ListDatasetsResponse, error) {
	resp := aiplatformpb.ListDatasetsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	resp.Datasets[0].SavedQueries[0].Metadata = nil
	return &resp, nil
}

func (*fakeDatasetLocationsRelationsServer) ListSavedQueries(context.Context, *aiplatformpb.ListSavedQueriesRequest) (*aiplatformpb.ListSavedQueriesResponse, error) {
	resp := aiplatformpb.ListSavedQueriesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
