package certificatemanager

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createCertificateIssuanceConfigs(gsrv *grpc.Server) error {
	fakeServer := &fakeCertificateIssuanceConfigsServer{}
	pb.RegisterCertificateManagerServer(gsrv, fakeServer)
	return nil
}

type fakeCertificateIssuanceConfigsServer struct {
	pb.UnimplementedCertificateManagerServer
}

func (*fakeCertificateIssuanceConfigsServer) ListCertificateIssuanceConfigs(context.Context, *pb.ListCertificateIssuanceConfigsRequest) (*pb.ListCertificateIssuanceConfigsResponse, error) {
	resp := pb.ListCertificateIssuanceConfigsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestCertificateIssuanceConfigs(t *testing.T) {
	client.MockTestGrpcHelper(t, CertificateIssuanceConfigs(), createCertificateIssuanceConfigs, client.TestOptions{})
}
