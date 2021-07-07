package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecurityPricings() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_pricings",
		Description:  "Pricing azure Security Center is provided in two pricing tiers: free and standard, with the standard tier available with a trial period The standard tier offers advanced security capabilities, while the free tier offers basic security features",
		Resolver:     fetchSecurityPricings,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "pricing_properties_tier",
				Description: "The pricing tier value Azure Security Center is provided in two pricing tiers: free and standard, with the standard tier available with a trial period The standard tier offers advanced security capabilities, while the free tier offers basic security features Possible values include: 'PricingTierFree', 'PricingTierStandard'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PricingProperties.PricingTier"),
			},
			{
				Name:        "pricing_properties_free_trial_remaining_time",
				Description: "The duration left for the subscriptions free trial period - in ISO 8601 format (eg P3Y6M4DT12H30M5S)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PricingProperties.FreeTrialRemainingTime"),
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "Resource type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Type"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSecurityPricings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Security.Pricings
	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	if response.Value != nil {
		for _, item := range *response.Value {
			res <- item
		}
	}
	return nil
}
