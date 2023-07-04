package batch

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/batch/apiv1/batchpb"
)

func createTaskGroups(gsrv *grpc.Server) error {
	fakeServer := &fakeTaskGroupsServer{}
	pb.RegisterBatchServiceServer(gsrv, fakeServer)
	return nil
}

type fakeTaskGroupsServer struct {
	pb.UnimplementedBatchServiceServer
}

func (*fakeTaskGroupsServer) ListJobs(context.Context, *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	resp := pb.ListJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeTaskGroupsServer) ListTasks(context.Context, *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	resp := pb.ListTasksResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestTaskGroups(t *testing.T) {
	client.MockTestGrpcHelper(t, TaskGroups(), createTaskGroups, client.TestOptions{})
}
