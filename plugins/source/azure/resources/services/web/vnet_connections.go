// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func vnetConnections() *schema.Table {
	return &schema.Table{
		Name:        "azure_web_vnet_connections",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#VnetInfo`,
		Resolver:    fetchWebVnetConnections,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "web_app_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "vnet_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VnetResourceID"),
			},
			{
				Name:     "cert_thumbprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertThumbprint"),
			},
			{
				Name:     "cert_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertBlob"),
			},
			{
				Name:     "routes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Routes"),
			},
			{
				Name:     "resync_required",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ResyncRequired"),
			},
			{
				Name:     "dns_servers",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSServers"),
			},
			{
				Name:     "is_swift",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsSwift"),
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchWebVnetConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.VnetConnections

	site := parent.Item.(web.Site)
	if site.SiteConfig == nil || site.SiteConfig.VnetName == nil {
		return nil
	}

	response, err := svc.GetVnetConnection(ctx, *site.ResourceGroup, *site.Name, *site.SiteConfig.VnetName)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
