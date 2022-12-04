// Code generated by codegen; DO NOT EDIT.

package armresources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TagDetails() *schema.Table {
	return &schema.Table{
		Name:      "azure_armresources_tag_details",
		Resolver:  fetchTagDetails,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "count",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Count"),
			},
			{
				Name:     "tag_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TagName"),
			},
			{
				Name:     "values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Values"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
		},
	}
}

func fetchTagDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armresources.NewTagsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
