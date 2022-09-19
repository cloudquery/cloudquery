// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func createSecurityPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNSecurityPoliciesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			SecurityPolicies: mockClient,
		},
	}

	data := cdn.SecurityPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewSecurityPolicyListResultPage(cdn.SecurityPolicyListResult{Value: &[]cdn.SecurityPolicy{data}}, func(ctx context.Context, result cdn.SecurityPolicyListResult) (cdn.SecurityPolicyListResult, error) {
		return cdn.SecurityPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
