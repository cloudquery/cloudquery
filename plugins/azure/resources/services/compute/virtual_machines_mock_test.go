//go:build !integration

package compute

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildComputeVirtualMachineMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	v := mocks.NewMockVirtualMachinesClient(ctrl)
	ve := mocks.NewMockVirtualMachineExtensionsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachines:          v,
			VirtualMachineExtensions: ve,
		},
	}
	virtualMachine := compute.VirtualMachine{}
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&virtualMachine)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	id := client.FakeResourceGroup + "/" + *virtualMachine.ID
	virtualMachine.ID = &id

	res := *virtualMachine.Resources
	res[0].Settings = "test"
	res[0].ProtectedSettings = "test"

	vmPage := compute.NewVirtualMachineListResultPage(compute.VirtualMachineListResult{Value: &[]compute.VirtualMachine{virtualMachine}}, func(ctx context.Context, result compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
		return compute.VirtualMachineListResult{}, nil
	})

	v.EXPECT().ListAll(gomock.Any(), gomock.Any()).Return(vmPage, nil)

	extension := compute.VirtualMachineExtension{}
	err = faker.FakeData(&extension)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	ve.EXPECT().
		List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(compute.VirtualMachineExtensionsListResult{Value: &[]compute.VirtualMachineExtension{extension}}, nil)

	instanceView := compute.VirtualMachineInstanceView{}
	err = faker.FakeData(&instanceView)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	v.EXPECT().
		InstanceView(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(instanceView, nil)

	return s
}

func TestComputeVirtualMachines(t *testing.T) {
	client.AzureMockTestHelper(t, ComputeVirtualMachines(), buildComputeVirtualMachineMock, client.TestOptions{})
}
