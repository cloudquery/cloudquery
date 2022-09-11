// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func TestComputeVirtualMachines(t *testing.T) {
	client.MockTestHelper(t, VirtualMachines(), createVirtualMachinesMock)
}

func createVirtualMachinesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachinesClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachines: mockClient,
		},
	}

	data := compute.VirtualMachine{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := compute.NewVirtualMachineListResultPage(compute.VirtualMachineListResult{Value: &[]compute.VirtualMachine{data}}, func(ctx context.Context, result compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
		return compute.VirtualMachineListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any(), "false").Return(result, nil)
	return s
}
