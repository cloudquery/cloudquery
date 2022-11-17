// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	compute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchVirtualMachineScaleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Compute

	virtualMachine := parent.Item.(*compute.VirtualMachine)
	id, err := arm.ParseResourceID(*virtualMachine.ID)
	if err != nil {
		return err
	}

	pager := svc.VirtualMachineScaleSetsClient.NewListPager(id.ResourceGroupName, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
