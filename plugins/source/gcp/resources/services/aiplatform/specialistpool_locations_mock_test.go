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

func createSpecialistpoolLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeSpecialistpoolLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeSpecialistpoolLocationsRelationsServer{}
	aiplatformpb.RegisterSpecialistPoolServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeSpecialistpoolLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (*fakeSpecialistpoolLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSpecialistpoolLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, SpecialistpoolLocations(), createSpecialistpoolLocations, client.TestOptions{})
}

type fakeSpecialistpoolLocationsRelationsServer struct {
	aiplatformpb.UnimplementedSpecialistPoolServiceServer
}

func (*fakeSpecialistpoolLocationsRelationsServer) ListSpecialistPools(context.Context, *aiplatformpb.ListSpecialistPoolsRequest) (*aiplatformpb.ListSpecialistPoolsResponse, error) {
	resp := aiplatformpb.ListSpecialistPoolsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
