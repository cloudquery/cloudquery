// Code generated by codegen; DO NOT EDIT.

package run

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "cloud.google.com/go/run/apiv2/runpb"
)

func createServices(gsrv *grpc.Server) error {
	fakeServer := &fakeServicesServer{}
	pb.RegisterServicesServer(gsrv, fakeServer)
	return nil
}

type fakeServicesServer struct {
	pb.UnimplementedServicesServer
}

func (f *fakeServicesServer) ListServices(context.Context, *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	resp := pb.ListServicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestServices(t *testing.T) {
	client.MockTestGrpcHelper(t, Services(), createServices, client.TestOptions{})
}
