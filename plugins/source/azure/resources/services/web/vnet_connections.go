// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func vnetConnections() *schema.Table {
	return &schema.Table{
		Name:     "azure_web_vnet_connections",
		Resolver: fetchWebVnetConnections,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
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
	response, err := svc.GetVnetConnection(ctx, *site.ResourceGroup, *site.Name, *site.SiteConfig.VnetName)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
