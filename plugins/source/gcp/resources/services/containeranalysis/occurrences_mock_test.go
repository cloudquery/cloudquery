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
	occ := pb.Occurrence{}
	if err := faker.FakeObject(&occ); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	occVul := pb.Occurrence_Vulnerability{}

	if err := faker.FakeObject(&occVul); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	occ.Details = &occVul
	return &pb.ListOccurrencesResponse{
		Occurrences: []*pb.Occurrence{&occ},
	}, nil
}

func TestOccurrences(t *testing.T) {
	client.MockTestGrpcHelper(t, Occurrences(), createOccurrences, client.TestOptions{})
}
