// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func SiteAuthSettings() *schema.Table {
	return &schema.Table{
		Name:     "azure_web_site_auth_settings",
		Resolver: fetchWebSiteAuthSettings,
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
				Name:     "site_auth_settings_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SiteAuthSettingsProperties"),
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

func fetchWebSiteAuthSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.SiteAuthSettings

	p := parent.Item.(web.Site)
	response, err := svc.GetAuthSettings(ctx, *p.ResourceGroup, *p.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
