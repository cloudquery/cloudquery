// Auto generated code - DO NOT EDIT.

package security

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Pricings() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_pricings",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security#Pricing`,
		Resolver:    fetchSecurityPricings,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "pricing_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PricingTier"),
			},
			{
				Name:     "free_trial_remaining_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FreeTrialRemainingTime"),
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

func fetchSecurityPricings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.Pricings

	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
