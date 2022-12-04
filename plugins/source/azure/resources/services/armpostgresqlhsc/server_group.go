// Code generated by codegen; DO NOT EDIT.

package armpostgresqlhsc

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServerGroup() *schema.Table {
	return &schema.Table{
		Name:      "azure_armpostgresqlhsc_server_group",
		Resolver:  fetchServerGroup,
		Multiplex: client.SubscriptionMultiplex,
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

func fetchServerGroup(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmpostgresqlhscServerGroup
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
