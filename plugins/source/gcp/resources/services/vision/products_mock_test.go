package vision

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
)

func createServer(gsrv *grpc.Server) error {
	fakeServer := &fakeServer{}
	pb.RegisterProductSearchServer(gsrv, fakeServer)
	return nil
}

type fakeServer struct {
	pb.UnimplementedProductSearchServer
}

func (*fakeServer) ListProducts(context.Context, *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	resp := pb.ListProductsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListReferenceImages(context.Context, *pb.ListReferenceImagesRequest) (*pb.ListReferenceImagesResponse, error) {
	resp := pb.ListReferenceImagesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, Products(), createServer, client.TestOptions{})
}
