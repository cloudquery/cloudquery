// Auto generated code - DO NOT EDIT.

package subscriptions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tenants() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscriptions_tenants",
		Resolver:  fetchSubscriptionsTenants,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "country",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Country"),
			},
			{
				Name:     "country_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CountryCode"),
			},
			{
				Name:     "default_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultDomain"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "domains",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Domains"),
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
				Name:     "tenant_branding_logo_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantBrandingLogoURL"),
			},
			{
				Name:     "tenant_category",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantCategory"),
			},
			{
				Name:     "tenant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantID"),
			},
			{
				Name:     "tenant_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantType"),
			},
		},
	}
}

func fetchSubscriptionsTenants(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions.Tenants
	pager := svc.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}
