package certificatemanager

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createCertificates(gsrv *grpc.Server) error {
	fakeServer := &fakeCertificatesServer{}
	pb.RegisterCertificateManagerServer(gsrv, fakeServer)
	return nil
}

type fakeCertificatesServer struct {
	pb.UnimplementedCertificateManagerServer
}

func (*fakeCertificatesServer) ListCertificates(context.Context, *pb.ListCertificatesRequest) (*pb.ListCertificatesResponse, error) {
	resp := pb.ListCertificatesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestCertificates(t *testing.T) {
	client.MockTestGrpcHelper(t, Certificates(), createCertificates, client.TestOptions{})
}
