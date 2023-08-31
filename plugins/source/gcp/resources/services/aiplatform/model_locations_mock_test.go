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

func createModelLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeModelLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeModelLocationsRelationsServer{}
	aiplatformpb.RegisterModelServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeModelLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeModelLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestModelLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, ModelLocations(), createModelLocations, client.TestOptions{})
}

type fakeModelLocationsRelationsServer struct {
	aiplatformpb.UnimplementedModelServiceServer
}

func (*fakeModelLocationsRelationsServer) ListModelEvaluationSlices(context.Context, *aiplatformpb.ListModelEvaluationSlicesRequest) (*aiplatformpb.ListModelEvaluationSlicesResponse, error) {
	resp := aiplatformpb.ListModelEvaluationSlicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeModelLocationsRelationsServer) ListModelEvaluations(context.Context, *aiplatformpb.ListModelEvaluationsRequest) (*aiplatformpb.ListModelEvaluationsResponse, error) {
	resp := aiplatformpb.ListModelEvaluationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeModelLocationsRelationsServer) ListModelVersions(context.Context, *aiplatformpb.ListModelVersionsRequest) (*aiplatformpb.ListModelVersionsResponse, error) {
	resp := aiplatformpb.ListModelVersionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeModelLocationsRelationsServer) ListModels(context.Context, *aiplatformpb.ListModelsRequest) (*aiplatformpb.ListModelsResponse, error) {
	resp := aiplatformpb.ListModelsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
