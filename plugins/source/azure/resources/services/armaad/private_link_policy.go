// Code generated by codegen; DO NOT EDIT.

package armaad

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PrivateLinkPolicy() *schema.Table {
	return &schema.Table{
		Name:      "azure_armaad_private_link_policy",
		Resolver:  fetchPrivateLinkPolicy,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "all_tenants",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllTenants"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owner_tenant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerTenantID"),
			},
			{
				Name:     "resource_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGroup"),
			},
			{
				Name:     "resource_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceName"),
			},
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionID"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "tenants",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tenants"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchPrivateLinkPolicy(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armaad.NewPrivateLinkForAzureAdClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
