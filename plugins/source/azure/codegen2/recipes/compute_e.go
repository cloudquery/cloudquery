package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	tables := []Table{
		{
			Service:        "armcompute",
			Name:           "skus",
			Struct:         &armcompute.ResourceSKU{},
			ResponseStruct: &armcompute.ResourceSKUsClientListResponse{},
			Client:         &armcompute.ResourceSKUsClient{},
			ListFunc:       (&armcompute.ResourceSKUsClient{}).NewListPager,
			NewFunc:        armcompute.NewResourceSKUsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/skus",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Name")`,
				},
			},
		},
		{
			Service:        "armcompute",
			Name:           "virtual_machines",
			Struct:         &armcompute.VirtualMachine{},
			ResponseStruct: &armcompute.VirtualMachinesClientListAllResponse{},
			Client:         &armcompute.VirtualMachinesClient{},
			ListFunc:       (&armcompute.VirtualMachinesClient{}).NewListAllPager,
			NewFunc:        armcompute.NewVirtualMachinesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			SkipMock:       true,
			ExtraColumns: []codegen.ColumnDefinition{
				SubscriptionIdColumn,
				{
					Name:     "instance_view",
					Type:     schema.TypeJSON,
					Resolver: `getInstanceView`,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}
