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

func createFirewallIngressRules(gsrv *grpc.Server) error {
	fakeServer := &fakeFirewallIngressRulesServer{}
	pb.RegisterFirewallServer(gsrv, fakeServer)
	return nil
}

type fakeFirewallIngressRulesServer struct {
	pb.UnimplementedFirewallServer
}

func (*fakeFirewallIngressRulesServer) ListIngressRules(context.Context, *pb.ListIngressRulesRequest) (*pb.ListIngressRulesResponse, error) {
	resp := pb.ListIngressRulesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestFirewallIngressRules(t *testing.T) {
	client.MockTestGrpcHelper(t, FirewallIngressRules(), createFirewallIngressRules, client.TestOptions{})
}
