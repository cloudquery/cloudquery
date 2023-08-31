package appengine

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAuthorizedDomains(gsrv *grpc.Server) error {
	fakeServer := &fakeAuthorizedDomainsServer{}
	pb.RegisterAuthorizedDomainsServer(gsrv, fakeServer)
	return nil
}

type fakeAuthorizedDomainsServer struct {
	pb.UnimplementedAuthorizedDomainsServer
}

func (*fakeAuthorizedDomainsServer) ListAuthorizedDomains(context.Context, *pb.ListAuthorizedDomainsRequest) (*pb.ListAuthorizedDomainsResponse, error) {
	resp := pb.ListAuthorizedDomainsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAuthorizedDomains(t *testing.T) {
	client.MockTestGrpcHelper(t, AuthorizedDomains(), createAuthorizedDomains, client.TestOptions{})
}
