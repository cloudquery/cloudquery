package monitoring

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
)

func createAlertPolicies(gsrv *grpc.Server) error {
	fakeServer := &fakeAlertPoliciesServer{}
	pb.RegisterAlertPolicyServiceServer(gsrv, fakeServer)
	return nil
}

type fakeAlertPoliciesServer struct {
	pb.UnimplementedAlertPolicyServiceServer
}

func (*fakeAlertPoliciesServer) ListAlertPolicies(context.Context, *pb.ListAlertPoliciesRequest) (*pb.ListAlertPoliciesResponse, error) {
	resp := pb.ListAlertPoliciesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAlertPolicies(t *testing.T) {
	client.MockTestGrpcHelper(t, AlertPolicies(), createAlertPolicies, client.TestOptions{})
}
