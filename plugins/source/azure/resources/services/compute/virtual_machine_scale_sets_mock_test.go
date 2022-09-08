// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func TestComputeVirtualMachineScaleSets(t *testing.T) {
	client.MockTestHelper(t, VirtualMachineScaleSets(), createVirtualMachineScaleSetsMock)
}

func createVirtualMachineScaleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachineScaleSetsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachineScaleSets: mockClient,
		},
	}

	data := compute.VirtualMachineScaleSet{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := compute.NewVirtualMachineScaleSetListWithLinkResultPage(compute.VirtualMachineScaleSetListWithLinkResult{Value: &[]compute.VirtualMachineScaleSet{data}}, func(ctx context.Context, result compute.VirtualMachineScaleSetListWithLinkResult) (compute.VirtualMachineScaleSetListWithLinkResult, error) {
		return compute.VirtualMachineScaleSetListWithLinkResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
