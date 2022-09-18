// Auto generated code - DO NOT EDIT.

package compute

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func createVirtualMachineExtensionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachineExtensionsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachineExtensions: mockClient,
		},
	}

	data := compute.VirtualMachineExtension{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.VirtualMachineExtensionsListResult{Value: &[]compute.VirtualMachineExtension{data}}

	mockClient.EXPECT().List(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
