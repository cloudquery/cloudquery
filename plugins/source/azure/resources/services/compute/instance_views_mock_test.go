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

func createInstanceViewsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeInstanceViewsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			InstanceViews: mockClient,
		},
	}

	data := compute.VirtualMachineInstanceView{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.VirtualMachineInstanceViewListResult{Value: &[]compute.VirtualMachineInstanceView{data}}

	mockClient.EXPECT().InstanceView(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
