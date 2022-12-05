// Code generated by codegen; DO NOT EDIT.

package armapplicationinsights

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ComponentApiKey() *schema.Table {
	return &schema.Table{
		Name:      "azure_armapplicationinsights_component_api_key",
		Resolver:  fetchComponentApiKey,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "created_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "linked_read_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedReadProperties"),
			},
			{
				Name:     "linked_write_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedWriteProperties"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "api_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIKey"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
		},
	}
}

func fetchComponentApiKey(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armapplicationinsights.NewAPIKeysClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
