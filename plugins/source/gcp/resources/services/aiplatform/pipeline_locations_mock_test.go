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

func createPipelineLocations(gsrv *grpc.Server) error {
	fakeServer := &fakePipelineLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakePipelineLocationsRelationsServer{}
	aiplatformpb.RegisterPipelineServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakePipelineLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakePipelineLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestPipelineLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, PipelineLocations(), createPipelineLocations, client.TestOptions{})
}

type fakePipelineLocationsRelationsServer struct {
	aiplatformpb.UnimplementedPipelineServiceServer
}

func (*fakePipelineLocationsRelationsServer) ListPipelineJobs(context.Context, *aiplatformpb.ListPipelineJobsRequest) (*aiplatformpb.ListPipelineJobsResponse, error) {
	resp := aiplatformpb.ListPipelineJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakePipelineLocationsRelationsServer) ListTrainingPipelines(context.Context, *aiplatformpb.ListTrainingPipelinesRequest) (*aiplatformpb.ListTrainingPipelinesResponse, error) {
	resp := aiplatformpb.ListTrainingPipelinesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
