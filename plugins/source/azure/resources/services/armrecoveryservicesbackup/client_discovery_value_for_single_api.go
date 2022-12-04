// Code generated by codegen; DO NOT EDIT.

package armrecoveryservicesbackup

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClientDiscoveryValueForSingleApi() *schema.Table {
	return &schema.Table{
		Name:      "azure_armrecoveryservicesbackup_client_discovery_value_for_single_api",
		Resolver:  fetchClientDiscoveryValueForSingleApi,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
		},
	}
}

func fetchClientDiscoveryValueForSingleApi(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmrecoveryservicesbackupClientDiscoveryValueForSingleApi
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
