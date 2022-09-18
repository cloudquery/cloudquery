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

func createCustomDomainsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNCustomDomainsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			CustomDomains: mockClient,
		},
	}

	data := cdn.CustomDomain{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewCustomDomainListResultPage(cdn.CustomDomainListResult{Value: &[]cdn.CustomDomain{data}}, func(ctx context.Context, result cdn.CustomDomainListResult) (cdn.CustomDomainListResult, error) {
		return cdn.CustomDomainListResult{}, nil
	})

	mockClient.EXPECT().ListByEndpoint(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
