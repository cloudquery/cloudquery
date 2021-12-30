package subscription

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SubscriptionSubscriptions() *schema.Table {
	return &schema.Table{
		Name:         "azure_subscription_subscriptions",
		Description:  "Model subscription information",
		Resolver:     fetchSubscriptionSubscriptions,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The fully qualified ID for the subscription For example, /subscriptions/00000000-0000-0000-0000-000000000000",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "subscription_id",
				Description: "The subscription ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionID"),
			},
			{
				Name:        "display_name",
				Description: "The subscription display name",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The subscription state Possible values are Enabled, Warned, PastDue, Disabled, and Deleted Possible values include: 'Enabled', 'Warned', 'PastDue', 'Disabled', 'Deleted'",
				Type:        schema.TypeString,
			},
			{
				Name:        "location_placement_id",
				Description: "The subscription location placement ID The ID indicates which regions are visible for a subscription For example, a subscription with a location placement Id of Public_2014-09-01 has access to Azure public regions",
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
				Description: "The subscription spending limit Possible values include: 'On', 'Off', 'CurrentPeriodOff'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.SpendingLimit"),
			},
			{
				Name:        "authorization_source",
				Description: "The authorization source of the request Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management For example, 'Legacy, RoleBased'",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSubscriptionSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	m, err := svc.Subscriptions.Get(ctx, svc.SubscriptionID)
	if err != nil {
		return err
	}
	res <- m
	return nil
}
