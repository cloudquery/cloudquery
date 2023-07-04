package certificatemanager

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
)

func createCertificateMaps(gsrv *grpc.Server) error {
	fakeServer := &fakeCertificateMapsServer{}
	pb.RegisterCertificateManagerServer(gsrv, fakeServer)
	return nil
}

type fakeCertificateMapsServer struct {
	pb.UnimplementedCertificateManagerServer
}

func (*fakeCertificateMapsServer) ListCertificateMaps(context.Context, *pb.ListCertificateMapsRequest) (*pb.ListCertificateMapsResponse, error) {
	resp := pb.ListCertificateMapsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeCertificateMapsServer) ListCertificateMapEntries(context.Context, *pb.ListCertificateMapEntriesRequest) (*pb.ListCertificateMapEntriesResponse, error) {
	resp := pb.ListCertificateMapEntriesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestCertificateMaps(t *testing.T) {
	client.MockTestGrpcHelper(t, CertificateMaps(), createCertificateMaps, client.TestOptions{})
}
