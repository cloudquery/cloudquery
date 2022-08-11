package subscription

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource subscriptions --config gen.hcl --output .
func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:         "azure_subscription_subscriptions",
		Description:  "Azure subscription information",
		Resolver:     fetchSubscriptionSubscriptions,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "authorization_source",
				Description: "The authorization source of the request",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by_tenants",
				Description: "An array containing the tenants managing the subscription",
				Type:        schema.TypeStringArray,
				Resolver:    resolveSubscriptionsManagedByTenants,
			},
			{
				Name:        "location_placement_id",
				Description: "The subscription location placement ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.LocationPlacementID"),
			},
			{
				Name:        "quota_id",
				Description: "The subscription quota ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.QuotaID"),
			},
			{
				Name:        "spending_limit",
				Description: "The subscription spending limit",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.SpendingLimit"),
			},
			{
				Name:          "tags",
				Description:   "The tags attached to the subscription",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "display_name",
				Description: "The subscription display name",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID for the subscription",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "state",
				Description: "The subscription state",
				Type:        schema.TypeString,
			},
			{
				Name:        "tenant_id",
				Description: "The subscription tenant ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TenantID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSubscriptionSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	pager := svc.Subscriptions.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}
func resolveSubscriptionsManagedByTenants(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(*armsubscriptions.Subscription)
	v := make([]*string, len(item.ManagedByTenants))
	for i, m := range item.ManagedByTenants {
		v[i] = m.TenantID
	}
	return diag.WrapError(resource.Set(c.Name, v))
}
