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

func createJobLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeJobLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeJobLocationsRelationsServer{}
	aiplatformpb.RegisterJobServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeJobLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeJobLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestJobLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, JobLocations(), createJobLocations, client.TestOptions{})
}

type fakeJobLocationsRelationsServer struct {
	aiplatformpb.UnimplementedJobServiceServer
}

func (*fakeJobLocationsRelationsServer) ListBatchPredictionJobs(context.Context, *aiplatformpb.ListBatchPredictionJobsRequest) (*aiplatformpb.ListBatchPredictionJobsResponse, error) {
	resp := aiplatformpb.ListBatchPredictionJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeJobLocationsRelationsServer) ListCustomJobs(context.Context, *aiplatformpb.ListCustomJobsRequest) (*aiplatformpb.ListCustomJobsResponse, error) {
	resp := aiplatformpb.ListCustomJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeJobLocationsRelationsServer) ListDataLabelingJobs(context.Context, *aiplatformpb.ListDataLabelingJobsRequest) (*aiplatformpb.ListDataLabelingJobsResponse, error) {
	resp := aiplatformpb.ListDataLabelingJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeJobLocationsRelationsServer) ListHyperparameterTuningJobs(context.Context, *aiplatformpb.ListHyperparameterTuningJobsRequest) (*aiplatformpb.ListHyperparameterTuningJobsResponse, error) {
	resp := aiplatformpb.ListHyperparameterTuningJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeJobLocationsRelationsServer) ListModelDeploymentMonitoringJobs(context.Context, *aiplatformpb.ListModelDeploymentMonitoringJobsRequest) (*aiplatformpb.ListModelDeploymentMonitoringJobsResponse, error) {
	resp := aiplatformpb.ListModelDeploymentMonitoringJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
