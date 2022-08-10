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

func buildComputeVirtualMachineScaleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	v := mocks.NewMockVirtualMachineScaleSetsClient(ctrl)

	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachineScaleSets: v,
		},
	}
	vmScaleSet := compute.VirtualMachineScaleSet{}
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&vmScaleSet)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	id := client.FakeResourceGroup + "/" + *vmScaleSet.ID
	vmScaleSet.ID = &id

	vmPage := compute.NewVirtualMachineScaleSetListWithLinkResultPage(compute.VirtualMachineScaleSetListWithLinkResult{Value: &[]compute.VirtualMachineScaleSet{vmScaleSet}},
		func(ctx context.Context, result compute.VirtualMachineScaleSetListWithLinkResult) (compute.VirtualMachineScaleSetListWithLinkResult, error) {
			return compute.VirtualMachineScaleSetListWithLinkResult{}, nil
		})

	v.EXPECT().ListAll(gomock.Any()).Return(vmPage, nil)

	return s
}

func TestComputeVirtualMachineScaleSets(t *testing.T) {
	client.AzureMockTestHelper(t, VirtualMachineScaleSets(), buildComputeVirtualMachineScaleSetsMock, client.TestOptions{})
}
