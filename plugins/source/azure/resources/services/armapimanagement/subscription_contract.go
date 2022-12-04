// Code generated by codegen; DO NOT EDIT.

package armapimanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubscriptionContract() *schema.Table {
	return &schema.Table{
		Name:      "azure_armapimanagement_subscription_contract",
		Resolver:  fetchSubscriptionContract,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSubscriptionContract(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armapimanagement.NewProductSubscriptionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
