package clouddeploy

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
)

func createDeliveryPipelines(gsrv *grpc.Server) error {
	fakeServer := &fakeDeliveryPipelinesServer{}
	pb.RegisterCloudDeployServer(gsrv, fakeServer)
	return nil
}

type fakeDeliveryPipelinesServer struct {
	pb.UnimplementedCloudDeployServer
}

func (*fakeDeliveryPipelinesServer) ListDeliveryPipelines(context.Context, *pb.ListDeliveryPipelinesRequest) (*pb.ListDeliveryPipelinesResponse, error) {
	resp := pb.ListDeliveryPipelinesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDeliveryPipelinesServer) ListReleases(context.Context, *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	resp := pb.ListReleasesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDeliveryPipelinesServer) ListRollouts(context.Context, *pb.ListRolloutsRequest) (*pb.ListRolloutsResponse, error) {
	resp := pb.ListRolloutsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDeliveryPipelinesServer) ListJobRuns(context.Context, *pb.ListJobRunsRequest) (*pb.ListJobRunsResponse, error) {
	resp := pb.ListJobRunsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestDeliveryPipelines(t *testing.T) {
	client.MockTestGrpcHelper(t, DeliveryPipelines(), createDeliveryPipelines, client.TestOptions{})
}
