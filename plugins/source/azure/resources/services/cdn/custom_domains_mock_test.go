// Code generated by codegen; DO NOT EDIT.

package cdn

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/cdn"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cdn"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCustomDomains(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	if c.Cdn == nil {
		c.Cdn = new(service.CdnClient)
	}
	cdnClient := c.Cdn
	if cdnClient.CustomDomainsClient == nil {
		cdnClient.CustomDomainsClient = mocks.NewMockCustomDomainsClient(ctrl)
	}

	mockCustomDomainsClient := cdnClient.CustomDomainsClient.(*mocks.MockCustomDomainsClient)

	var response api.CustomDomainsClientListByEndpointResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockCustomDomainsClient.EXPECT().NewListByEndpointPager(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)
}
