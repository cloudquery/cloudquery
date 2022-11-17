// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/compute"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/compute"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildVirtualMachineScaleSets(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	if c.Compute == nil {
		c.Compute = new(service.ComputeClient)
	}
	computeClient := c.Compute
	if computeClient.VirtualMachineScaleSetsClient == nil {
		computeClient.VirtualMachineScaleSetsClient = mocks.NewMockVirtualMachineScaleSetsClient(ctrl)
	}

	mockVirtualMachineScaleSetsClient := computeClient.VirtualMachineScaleSetsClient.(*mocks.MockVirtualMachineScaleSetsClient)

	var response api.VirtualMachineScaleSetsClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockVirtualMachineScaleSetsClient.EXPECT().NewListPager(gomock.Any(), gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)
}
