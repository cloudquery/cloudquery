// Code generated by codegen; DO NOT EDIT.

package armnetapp

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CapacityPool() *schema.Table {
	return &schema.Table{
		Name:      "azure_armnetapp_capacity_pool",
		Resolver:  fetchCapacityPool,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchCapacityPool(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armnetapp.NewPoolsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
