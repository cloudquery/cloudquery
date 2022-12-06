// Code generated by codegen; DO NOT EDIT.

package subscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Subscription() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscription_subscription",
		Resolver:  fetchSubscription,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "authorization_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizationSource"),
			},
			{
				Name:     "subscription_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SubscriptionPolicies"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
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
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionID"),
			},
		},
	}
}

func fetchSubscription(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armsubscription.NewSubscriptionsClient(cl.Creds, cl.Options)
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
