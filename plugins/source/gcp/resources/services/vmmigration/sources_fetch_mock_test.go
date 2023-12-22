package vmmigration

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
)

func createSourcesServer(gsrv *grpc.Server) error {
	fakeServer := &fakeSourcesServer{}
	pb.RegisterVmMigrationServer(gsrv, fakeServer)
	return nil
}

type fakeSourcesServer struct {
	pb.UnimplementedVmMigrationServer
}

func (*fakeSourcesServer) ListSources(context.Context, *pb.ListSourcesRequest) (*pb.ListSourcesResponse, error) {
	resp := pb.ListSourcesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.Sources[0].SourceDetails = &pb.Source_Aws{Aws: &pb.AwsSourceDetails{AwsRegion: "us-east-1"}}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeSourcesServer) ListMigratingVms(context.Context, *pb.ListMigratingVmsRequest) (*pb.ListMigratingVmsResponse, error) {
	resp := pb.ListMigratingVmsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.MigratingVms[0].TargetVmDefaults = &pb.MigratingVm_ComputeEngineTargetDefaults{ComputeEngineTargetDefaults: &pb.ComputeEngineTargetDefaults{VmName: "test-vm"}}
	resp.MigratingVms[0].SourceVmDetails = &pb.MigratingVm_AwsSourceVmDetails{AwsSourceVmDetails: &pb.AwsSourceVmDetails{CommittedStorageBytes: 1}}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeSourcesServer) ListDatacenterConnectors(context.Context, *pb.ListDatacenterConnectorsRequest) (*pb.ListDatacenterConnectorsResponse, error) {
	resp := pb.ListDatacenterConnectorsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeSourcesServer) ListCloneJobs(context.Context, *pb.ListCloneJobsRequest) (*pb.ListCloneJobsResponse, error) {
	resp := pb.ListCloneJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.CloneJobs[0].TargetVmDetails = &pb.CloneJob_ComputeEngineTargetDetails{ComputeEngineTargetDetails: &pb.ComputeEngineTargetDetails{VmName: "test-vm"}}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeSourcesServer) ListCutoverJobs(context.Context, *pb.ListCutoverJobsRequest) (*pb.ListCutoverJobsResponse, error) {
	resp := pb.ListCutoverJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.CutoverJobs[0].TargetVmDetails = &pb.CutoverJob_ComputeEngineTargetDetails{ComputeEngineTargetDetails: &pb.ComputeEngineTargetDetails{VmName: "test-vm"}}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeSourcesServer) ListUtilizationReports(context.Context, *pb.ListUtilizationReportsRequest) (*pb.ListUtilizationReportsResponse, error) {
	resp := pb.ListUtilizationReportsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestHelper(t, Sources(), client.WithCreateGrpcService(createSourcesServer))
}
