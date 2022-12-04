// Code generated by codegen; DO NOT EDIT.

package armsubscriptions

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TenantIdDescription() *schema.Table {
	return &schema.Table{
		Name:      "azure_armsubscriptions_tenant_id_description",
		Resolver:  fetchTenantIdDescription,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
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

func fetchTenantIdDescription(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armsubscriptions.NewTenantsClient(cl.Creds, cl.Options)
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
