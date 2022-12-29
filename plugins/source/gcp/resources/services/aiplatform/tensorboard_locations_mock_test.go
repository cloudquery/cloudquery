// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "google.golang.org/genproto/googleapis/cloud/location"

	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
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

func (f *fakeTensorboardLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
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

func (f *fakeTensorboardLocationsRelationsServer) ListTensorboardExperiments(context.Context, *aiplatformpb.ListTensorboardExperimentsRequest) (*aiplatformpb.ListTensorboardExperimentsResponse, error) {
	resp := aiplatformpb.ListTensorboardExperimentsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeTensorboardLocationsRelationsServer) ListTensorboardRuns(context.Context, *aiplatformpb.ListTensorboardRunsRequest) (*aiplatformpb.ListTensorboardRunsResponse, error) {
	resp := aiplatformpb.ListTensorboardRunsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeTensorboardLocationsRelationsServer) ListTensorboardTimeSeries(context.Context, *aiplatformpb.ListTensorboardTimeSeriesRequest) (*aiplatformpb.ListTensorboardTimeSeriesResponse, error) {
	resp := aiplatformpb.ListTensorboardTimeSeriesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeTensorboardLocationsRelationsServer) ListTensorboards(context.Context, *aiplatformpb.ListTensorboardsRequest) (*aiplatformpb.ListTensorboardsResponse, error) {
	resp := aiplatformpb.ListTensorboardsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
