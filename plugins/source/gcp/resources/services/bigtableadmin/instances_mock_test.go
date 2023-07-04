package bigtableadmin

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func createInstances(gsrv *grpc.Server) error {
	pb.RegisterBigtableInstanceAdminServer(gsrv, &fakeInstanceAdminServer{})
	pb.RegisterBigtableTableAdminServer(gsrv, &fakeTableServer{})
	return nil
}

type fakeInstanceAdminServer struct {
	pb.UnimplementedBigtableInstanceAdminServer
}

func (*fakeInstanceAdminServer) ListInstances(context.Context, *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	resp := pb.ListInstancesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.Instances[0].Name = "projects/testProject/instances/test-instance"
	resp.FailedLocations = nil
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeInstanceAdminServer) ListAppProfiles(context.Context, *pb.ListAppProfilesRequest) (*pb.ListAppProfilesResponse, error) {
	resp := pb.ListAppProfilesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeInstanceAdminServer) ListClusters(context.Context, *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	resp := pb.ListClustersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	clusterConfig := pb.Cluster_ClusterConfig_{}
	if err := faker.FakeObject(&clusterConfig); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.Clusters[0].Config = &clusterConfig
	resp.FailedLocations = nil
	resp.NextPageToken = ""
	return &resp, nil
}

type fakeTableServer struct {
	pb.UnimplementedBigtableTableAdminServer
}

func (*fakeTableServer) ListTables(context.Context, *pb.ListTablesRequest) (*pb.ListTablesResponse, error) {
	resp := pb.ListTablesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeTableServer) GetTable(context.Context, *pb.GetTableRequest) (*pb.Table, error) {
	resp := pb.Table{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func (*fakeTableServer) ListBackups(context.Context, *pb.ListBackupsRequest) (*pb.ListBackupsResponse, error) {
	resp := pb.ListBackupsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Instances(), createInstances, client.TestOptions{})
}
