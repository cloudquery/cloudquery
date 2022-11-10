// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func siteAuthSettingsV2() *schema.Table {
	return &schema.Table{
		Name:        "azure_web_site_auth_settings_v2",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#SiteAuthSettingsV2`,
		Resolver:    fetchWebSiteAuthSettingsV2,
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
				Name:     "platform",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "global_validation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalValidation"),
			},
			{
				Name:     "identity_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviders"),
			},
			{
				Name:     "login",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Login"),
			},
			{
				Name:     "http_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HTTPSettings"),
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

func fetchWebSiteAuthSettingsV2(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.SiteAuthSettingsV2

	site := parent.Item.(web.Site)
	response, err := svc.GetAuthSettingsV2(ctx, *site.ResourceGroup, *site.Name)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
