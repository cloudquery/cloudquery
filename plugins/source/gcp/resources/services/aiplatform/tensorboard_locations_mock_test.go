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

func createTensorboardLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeTensorboardLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeTensorboardLocationsRelationsServer{}
	aiplatformpb.RegisterTensorboardServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeTensorboardLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeTensorboardLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestTensorboardLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, TensorboardLocations(), createTensorboardLocations, client.TestOptions{})
}

type fakeTensorboardLocationsRelationsServer struct {
	aiplatformpb.UnimplementedTensorboardServiceServer
}

func (*fakeTensorboardLocationsRelationsServer) ListTensorboardExperiments(context.Context, *aiplatformpb.ListTensorboardExperimentsRequest) (*aiplatformpb.ListTensorboardExperimentsResponse, error) {
	resp := aiplatformpb.ListTensorboardExperimentsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeTensorboardLocationsRelationsServer) ListTensorboardRuns(context.Context, *aiplatformpb.ListTensorboardRunsRequest) (*aiplatformpb.ListTensorboardRunsResponse, error) {
	resp := aiplatformpb.ListTensorboardRunsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeTensorboardLocationsRelationsServer) ListTensorboardTimeSeries(context.Context, *aiplatformpb.ListTensorboardTimeSeriesRequest) (*aiplatformpb.ListTensorboardTimeSeriesResponse, error) {
	resp := aiplatformpb.ListTensorboardTimeSeriesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeTensorboardLocationsRelationsServer) ListTensorboards(context.Context, *aiplatformpb.ListTensorboardsRequest) (*aiplatformpb.ListTensorboardsResponse, error) {
	resp := aiplatformpb.ListTensorboardsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
