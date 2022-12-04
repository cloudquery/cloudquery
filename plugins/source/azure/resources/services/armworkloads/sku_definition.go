// Code generated by codegen; DO NOT EDIT.

package armworkloads

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SkuDefinition() *schema.Table {
	return &schema.Table{
		Name:      "azure_armworkloads_sku_definition",
		Resolver:  fetchSkuDefinition,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capabilities"),
			},
			{
				Name:     "costs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Costs"),
			},
			{
				Name:     "family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Family"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "location_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocationInfo"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "restrictions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Restrictions"),
			},
			{
				Name:     "size",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
		},
	}
}

func fetchSkuDefinition(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmworkloadsSkuDefinition
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
