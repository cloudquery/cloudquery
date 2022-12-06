// Code generated by codegen; DO NOT EDIT.

package dns

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Zone() *schema.Table {
	return &schema.Table{
		Name:      "azure_dns_zone",
		Resolver:  fetchZone,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
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

func fetchZone(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdns.NewZonesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
