package containeranalysis

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createOccurrences(gsrv *grpc.Server) error {
	fakeServer := &fakeOccurrencesServer{}
	pb.RegisterGrafeasV1Beta1Server(gsrv, fakeServer)
	return nil
}

type fakeOccurrencesServer struct {
	pb.UnimplementedGrafeasV1Beta1Server
}

func (*fakeOccurrencesServer) ListOccurrences(context.Context, *pb.ListOccurrencesRequest) (*pb.ListOccurrencesResponse, error) {
	resp := pb.ListOccurrencesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestOccurrences(t *testing.T) {
	client.MockTestGrpcHelper(t, Occurrences(), createOccurrences, client.TestOptions{})
}
