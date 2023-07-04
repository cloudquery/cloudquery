package billing

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	budgetspb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
)

func createBillingAccounts(gsrv *grpc.Server) error {
	pb.RegisterCloudBillingServer(gsrv, &fakeBillingAccountsServer{})
	budgetspb.RegisterBudgetServiceServer(gsrv, &fakeBudgetsServer{})
	return nil
}

type fakeBillingAccountsServer struct {
	pb.UnimplementedCloudBillingServer
}

func (*fakeBillingAccountsServer) ListBillingAccounts(context.Context, *pb.ListBillingAccountsRequest) (*pb.ListBillingAccountsResponse, error) {
	resp := pb.ListBillingAccountsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

type fakeBudgetsServer struct {
	budgetspb.UnimplementedBudgetServiceServer
}

func (*fakeBudgetsServer) ListBudgets(context.Context, *budgetspb.ListBudgetsRequest) (*budgetspb.ListBudgetsResponse, error) {
	resp := budgetspb.ListBudgetsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestBillingAccounts(t *testing.T) {
	client.MockTestGrpcHelper(t, BillingAccounts(), createBillingAccounts, client.TestOptions{})
}
