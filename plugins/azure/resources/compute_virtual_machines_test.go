package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildComputeVirtualMachineMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	v := mocks.NewMockVirtualMachinesClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachines: v,
		},
	}
	virtualMachine := compute.VirtualMachine{}
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&virtualMachine)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	res := *virtualMachine.Resources
	res[0].Settings = "test"
	res[0].ProtectedSettings = "test"

	vmPage := compute.NewVirtualMachineListResultPage(compute.VirtualMachineListResult{Value: &[]compute.VirtualMachine{virtualMachine}}, func(ctx context.Context, result compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
		//return result, nil
		return compute.VirtualMachineListResult{}, nil
	})

	v.EXPECT().ListAll(gomock.Any(), gomock.Any()).Return(vmPage, nil)

	return s
}

func TestComputeVirtualMachines(t *testing.T) {
	azureTestHelper(t, resources.ComputeVirtualMachines(), buildComputeVirtualMachineMock)
}
