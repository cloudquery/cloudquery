// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func customDomains() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_custom_domains",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn#CustomDomain`,
		Resolver:    fetchCDNCustomDomains,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cdn_endpoint_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostName"),
			},
			{
				Name:     "resource_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceState"),
			},
			{
				Name:     "custom_https_provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomHTTPSProvisioningState"),
			},
			{
				Name:     "custom_https_provisioning_substate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomHTTPSProvisioningSubstate"),
			},
			{
				Name:     "validation_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValidationData"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
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
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
		},
	}
}

func fetchCDNCustomDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CDN.CustomDomains

	profile := parent.Parent.Item.(cdn.Profile)
	resource, err := client.ParseResourceID(*profile.ID)
	if err != nil {
		return err
	}
	endpoint := parent.Item.(cdn.Endpoint)
	response, err := svc.ListByEndpoint(ctx, resource.ResourceGroup, *profile.Name, *endpoint.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
