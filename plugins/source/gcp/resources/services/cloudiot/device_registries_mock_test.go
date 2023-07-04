package cloudiot

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/iot/apiv1/iotpb"
)

func createDeviceRegistries(gsrv *grpc.Server) error {
	fakeServer := &fakeDeviceRegistriesServer{}
	pb.RegisterDeviceManagerServer(gsrv, fakeServer)
	return nil
}

type fakeDeviceRegistriesServer struct {
	pb.UnimplementedDeviceManagerServer
}

func (*fakeDeviceRegistriesServer) ListDeviceRegistries(context.Context, *pb.ListDeviceRegistriesRequest) (*pb.ListDeviceRegistriesResponse, error) {
	resp := pb.ListDeviceRegistriesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDeviceRegistriesServer) ListDevices(context.Context, *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {
	resp := pb.ListDevicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeDeviceRegistriesServer) ListDeviceConfigVersions(context.Context, *pb.ListDeviceConfigVersionsRequest) (*pb.ListDeviceConfigVersionsResponse, error) {
	resp := pb.ListDeviceConfigVersionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func (*fakeDeviceRegistriesServer) ListDeviceStates(context.Context, *pb.ListDeviceStatesRequest) (*pb.ListDeviceStatesResponse, error) {
	resp := pb.ListDeviceStatesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestDeviceRegistries(t *testing.T) {
	client.MockTestGrpcHelper(t, DeviceRegistries(), createDeviceRegistries, client.TestOptions{})
}
