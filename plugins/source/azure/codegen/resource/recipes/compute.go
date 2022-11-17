package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Compute() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armcompute.Disk),
			Resolver: &resource.FuncParams{
				Func: compute.DisksClient.NewListPager,
			},
		},
		{
			Struct: new(armcompute.VirtualMachine),
			Resolver: &resource.FuncParams{
				Func: compute.VirtualMachinesClient.NewListAllPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armcompute.VirtualMachineScaleSet),
					Resolver: &resource.FuncParams{
						Func:   compute.VirtualMachineScaleSetsClient.NewListPager,
						Params: []string{"id.ResourceGroupName"},
					},
				},
			},
		},
	}
}
