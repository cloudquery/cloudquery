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
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Compute)`,
			ExtraColumns:   []codegen.ColumnDefinition{
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Name")`,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}
