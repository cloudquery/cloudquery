// Code generated by codegen; DO NOT EDIT.

package armelasticsan

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SkuInformation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armelasticsan_sku_information",
		Resolver:  fetchSkuInformation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capabilities"),
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
		},
	}
}

func fetchSkuInformation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmelasticsanSkuInformation
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
