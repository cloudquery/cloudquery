// Code generated by codegen; DO NOT EDIT.

package notificationhubs

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func NamespaceResource() *schema.Table {
	return &schema.Table{
		Name:      "azure_notificationhubs_namespace_resource",
		Resolver:  fetchNamespaceResource,
		Multiplex: client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.NotificationHubs"),
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
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SKU"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchNamespaceResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armnotificationhubs.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
