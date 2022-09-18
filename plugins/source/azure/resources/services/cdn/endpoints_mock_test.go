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

func createEndpointsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNEndpointsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Endpoints:     mockClient,
			CustomDomains: createCustomDomainsMock(t, ctrl).CDN.CustomDomains,
			Routes:        createRoutesMock(t, ctrl).CDN.Routes,
		},
	}

	data := cdn.Endpoint{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := cdn.NewEndpointListResultPage(cdn.EndpointListResult{Value: &[]cdn.Endpoint{data}}, func(ctx context.Context, result cdn.EndpointListResult) (cdn.EndpointListResult, error) {
		return cdn.EndpointListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
