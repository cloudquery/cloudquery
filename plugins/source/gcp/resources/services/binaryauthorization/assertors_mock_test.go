package binaryauthorization

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAssertors(gsrv *grpc.Server) error {
	fakeServer := &fakeAssertorsServer{}
	pb.RegisterBinauthzManagementServiceV1Server(gsrv, fakeServer)
	return nil
}

type fakeAssertorsServer struct {
	pb.UnimplementedBinauthzManagementServiceV1Server
}

func (*fakeAssertorsServer) ListAttestors(context.Context, *pb.ListAttestorsRequest) (*pb.ListAttestorsResponse, error) {
	resp := pb.ListAttestorsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAssertors(t *testing.T) {
	client.MockTestGrpcHelper(t, Assertors(), createAssertors, client.TestOptions{})
}
