package security

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecurityJitNetworkAccessPolicies(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockJitNetworkAccessPoliciesClient(ctrl)

	var policy security.JitNetworkAccessPolicy
	if err := faker.FakeData(&policy); err != nil {
		t.Fatal(err)
	}

	id := client.FakeResourceGroup + "/" + *policy.ID
	policy.ID = &id
	result := security.NewJitNetworkAccessPoliciesListPage(
		security.JitNetworkAccessPoliciesList{Value: &[]security.JitNetworkAccessPolicy{policy}},
		func(context.Context, security.JitNetworkAccessPoliciesList) (security.JitNetworkAccessPoliciesList, error) {
			return security.JitNetworkAccessPoliciesList{}, nil
		},
	)
	ip := faker.IPv4()
	(*policy.VirtualMachines)[0].PublicIPAddress = &ip
	m.EXPECT().List(gomock.Any()).Return(result, nil)
	return services.Services{
		Security: services.SecurityClient{JitNetworkAccessPolicies: m},
	}
}

func TestSecurityJitNetworkAccessPolicies(t *testing.T) {
	client.AzureMockTestHelper(t, SecurityJitNetworkAccessPolicies(), buildSecurityJitNetworkAccessPolicies, client.TestOptions{})
}
