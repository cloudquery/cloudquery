// Code generated by codegen; DO NOT EDIT.

package armstoragecache

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceUsage() *schema.Table {
	return &schema.Table{
		Name:      "azure_armstoragecache_resource_usage",
		Resolver:  fetchResourceUsage,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "current_value",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CurrentValue"),
			},
			{
				Name:     "limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Limit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
		},
	}
}

func fetchResourceUsage(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmstoragecacheResourceUsage
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
