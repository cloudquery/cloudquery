package apikeys

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/genproto/googleapis/api/apikeys/v2"
	"google.golang.org/grpc"
)

func createKeys(gsrv *grpc.Server) error {
	fakeServer := &fakeKeysServer{}
	pb.RegisterApiKeysServer(gsrv, fakeServer)
	return nil
}

type fakeKeysServer struct {
	pb.UnimplementedApiKeysServer
}

func (*fakeKeysServer) ListKeys(context.Context, *pb.ListKeysRequest) (*pb.ListKeysResponse, error) {
	resp := pb.ListKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestKeys(t *testing.T) {
	client.MockTestGrpcHelper(t, Keys(), createKeys, client.TestOptions{})
}
