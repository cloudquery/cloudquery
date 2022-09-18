// Auto generated code - DO NOT EDIT.

package subscriptions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscriptions_subscriptions",
		Resolver:  fetchSubscriptionsSubscriptions,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "authorization_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizationSource"),
			},
			{
				Name:     "managed_by_tenants",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ManagedByTenants"),
			},
			{
				Name:     "subscription_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SubscriptionPolicies"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
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
			{
				Name:     "tenant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantID"),
			},
		},
	}
}

func fetchSubscriptionsSubscriptions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions.Subscriptions
	pager := svc.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}
