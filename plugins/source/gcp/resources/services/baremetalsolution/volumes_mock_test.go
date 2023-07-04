package baremetalsolution

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
)

func createVolumes(gsrv *grpc.Server) error {
	pb.RegisterBareMetalSolutionServer(gsrv, &fakeVolumesServer{})
	return nil
}

type fakeVolumesServer struct {
	pb.UnimplementedBareMetalSolutionServer
}

func (*fakeVolumesServer) ListVolumes(context.Context, *pb.ListVolumesRequest) (*pb.ListVolumesResponse, error) {
	resp := pb.ListVolumesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeVolumesServer) ListLuns(context.Context, *pb.ListLunsRequest) (*pb.ListLunsResponse, error) {
	resp := pb.ListLunsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestVolumes(t *testing.T) {
	client.MockTestGrpcHelper(t, Volumes(), createVolumes, client.TestOptions{})
}
