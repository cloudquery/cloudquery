// Code generated by codegen; DO NOT EDIT.

package armstoragepool

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceSkuInfo() *schema.Table {
	return &schema.Table{
		Name:      "azure_armstoragepool_resource_sku_info",
		Resolver:  fetchResourceSkuInfo,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
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
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
		},
	}
}

func fetchResourceSkuInfo(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmstoragepoolResourceSkuInfo
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
