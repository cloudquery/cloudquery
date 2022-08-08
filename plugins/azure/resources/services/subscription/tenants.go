package subscription

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource tenants --config gen.hcl --output .
func Tenants() *schema.Table {
	return &schema.Table{
		Name:        "azure_subscription_tenants",
		Description: "Azure tenant information",
		Resolver:    fetchSubscriptionTenants,
		Multiplex:   client.SubscriptionMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:          "country",
				Description:   "Country/region name of the address for the tenant",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "country_code",
				Description:   "Country/region abbreviation for the tenant",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "default_domain",
				Description:   "The default domain for the tenant",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "display_name",
				Description:   "The display name of the tenant",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "domains",
				Description:   "The list of domains for the tenant",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID of the tenant",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:          "tenant_branding_logo_url",
				Description:   "The tenant's branding logo URL",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("TenantBrandingLogoURL"),
				IgnoreInTests: true,
			},
			{
				Name:        "tenant_category",
				Description: "Category of the tenant",
				Type:        schema.TypeString,
			},
			{
				Name:        "tenant_id",
				Description: "The tenant ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TenantID"),
			},
			{
				Name:          "tenant_type",
				Description:   "The tenant type",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSubscriptionTenants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	pager := svc.Tenants.NewListPager(nil)
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
