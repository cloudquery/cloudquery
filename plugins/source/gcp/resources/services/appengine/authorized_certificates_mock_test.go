package appengine

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
)

func createAuthorizedCertificates(gsrv *grpc.Server) error {
	fakeServer := &fakeAuthorizedCertificatesServer{}
	pb.RegisterAuthorizedCertificatesServer(gsrv, fakeServer)
	return nil
}

type fakeAuthorizedCertificatesServer struct {
	pb.UnimplementedAuthorizedCertificatesServer
}

func (*fakeAuthorizedCertificatesServer) ListAuthorizedCertificates(context.Context, *pb.ListAuthorizedCertificatesRequest) (*pb.ListAuthorizedCertificatesResponse, error) {
	resp := pb.ListAuthorizedCertificatesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAuthorizedCertificates(t *testing.T) {
	client.MockTestGrpcHelper(t, AuthorizedCertificates(), createAuthorizedCertificates, client.TestOptions{})
}
